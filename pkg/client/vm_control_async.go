// Copyright (c) ZStack.io, Inc.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// VMControlResult is an API observation, not physical-release evidence.
type VMControlResult struct{ State, JobID string }

// ControlVmInstanceAsync sends one bounded SDK request and does not retry. The
// stable job UUID is supplied by the durable resource owner before submission.
func (cli *ZSClient) ControlVmInstanceAsync(ctx context.Context, job, vm, action string) (VMControlResult, error) {
	if cli.readOnly || !asyncID(job) || !asyncID(vm) {
		return VMControlResult{}, ErrAsyncRequest
	}
	path, method := "v1/vm-instances/"+vm+"/actions", http.MethodPut
	var params any
	switch action {
	case "stop":
		params = map[string]any{"stopVmInstance": map[string]any{"type": "grace"}}
	case "force-stop":
		params = map[string]any{"stopVmInstance": map[string]any{"type": "cold"}}
	case "destroy":
		method = http.MethodDelete
		path = "v1/vm-instances/" + vm + "?deleteMode=Enforcing"
	case "expunge":
		params = map[string]any{"expungeVmInstance": map[string]any{}}
	default:
		return VMControlResult{}, ErrAsyncRequest
	}
	var body []byte
	var e error
	if params != nil {
		body, e = json.Marshal(params)
		if e != nil {
			return VMControlResult{}, ErrAsyncRequest
		}
	}
	return cli.vmControlRequest(ctx, job, vm, path, method, body, action == "destroy" || action == "expunge")
}
func (cli *ZSClient) QueryVmControl(ctx context.Context, job, vm string) (VMControlResult, error) {
	if !asyncID(job) || !asyncID(vm) {
		return VMControlResult{}, ErrAsyncRequest
	}
	return cli.vmControlRequest(ctx, job, vm, "v1/api-jobs/"+job, http.MethodGet, nil, false)
}

// QueryVmControlForAction accepts Cloud's canonical empty-object receipt only
// for a persisted destroy/expunge action. It is not physical-release evidence.
func (cli *ZSClient) QueryVmControlForAction(ctx context.Context, job, vm, action string) (VMControlResult, error) {
	if !asyncID(job) || !asyncID(vm) || (action != "stop" && action != "force-stop" && action != "destroy" && action != "expunge") {
		return VMControlResult{}, ErrAsyncRequest
	}
	return cli.vmControlRequest(ctx, job, vm, "v1/api-jobs/"+job, http.MethodGet, nil, action == "destroy" || action == "expunge")
}
func (cli *ZSClient) vmControlRequest(ctx context.Context, job, vm, path, method string, body []byte, emptyReceipt bool) (VMControlResult, error) {
	target := cli.getRequestURL(path)
	header, e := cli.getHeader(target, method)
	if e != nil {
		return VMControlResult{}, ErrAsyncRequest
	}
	req, e := http.NewRequestWithContext(ctx, method, target, bytes.NewReader(body))
	if e != nil {
		return VMControlResult{}, ErrAsyncRequest
	}
	req.Header = header
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if method != http.MethodGet {
		req.Header.Set("X-Job-UUID", job)
	}
	hc := *cli.httpClient
	hc.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	response, e := hc.Do(req)
	if e != nil {
		if ctx.Err() != nil {
			return VMControlResult{}, ctx.Err()
		}
		return VMControlResult{}, ErrAsyncRequest
	}
	defer response.Body.Close()
	raw, e := io.ReadAll(io.LimitReader(response.Body, (2<<20)+1))
	if e != nil || len(raw) > 2<<20 {
		return VMControlResult{}, ErrAsyncResponse
	}
	if response.StatusCode == 404 && method == http.MethodGet {
		return VMControlResult{}, ErrAsyncNotFound
	}
	if response.StatusCode == 204 && method == http.MethodDelete && len(bytes.TrimSpace(raw)) == 0 {
		return VMControlResult{State: "accepted", JobID: job}, nil
	}
	if response.StatusCode != 200 && response.StatusCode != 202 && response.StatusCode != 503 {
		return VMControlResult{}, ErrAsyncResponse
	}
	if response.StatusCode == 202 && method == http.MethodGet && len(bytes.TrimSpace(raw)) == 0 {
		return VMControlResult{State: "running", JobID: job}, nil
	}
	var v map[string]json.RawMessage
	if json.Unmarshal(raw, &v) != nil || v == nil {
		return VMControlResult{}, ErrAsyncResponse
	}
	if b, ok := v["error"]; ok {
		var err struct {
			Code string `json:"code"`
		}
		if json.Unmarshal(b, &err) != nil || err.Code == "" || (response.StatusCode != 200 && response.StatusCode != 503) {
			return VMControlResult{}, ErrAsyncResponse
		}
		return VMControlResult{State: "failed", JobID: job}, nil
	}
	if response.StatusCode == 503 {
		return VMControlResult{}, ErrAsyncResponse
	}
	if b, ok := v["success"]; ok && string(b) != "true" {
		return VMControlResult{}, ErrAsyncResponse
	}
	if response.StatusCode == 202 {
		if b, ok := v["location"]; ok {
			var location string
			if json.Unmarshal(b, &location) != nil {
				return VMControlResult{}, ErrAsyncResponse
			}
			id, e := cli.asyncJobID(location)
			if e != nil || id != job {
				return VMControlResult{}, ErrAsyncResponse
			}
		} else if method != http.MethodGet {
			return VMControlResult{}, ErrAsyncResponse
		}
		state := "running"
		if method != http.MethodGet {
			state = "accepted"
		}
		return VMControlResult{State: state, JobID: job}, nil
	}
	if _, ok := v["location"]; ok {
		return VMControlResult{}, ErrAsyncResponse
	}
	// APIDestroy/ExpungeVmInstanceEvent have @RestResponse with no fields;
	// a completed successful job is HTTP 200 {}. Other actions remain strict.
	if emptyReceipt && len(v) == 0 {
		return VMControlResult{State: "succeeded", JobID: job}, nil
	}
	if b, ok := v["inventory"]; ok {
		var inventory struct {
			UUID string `json:"uuid"`
		}
		if json.Unmarshal(b, &inventory) != nil || inventory.UUID != vm {
			return VMControlResult{}, ErrAsyncResponse
		}
		return VMControlResult{State: "succeeded", JobID: job}, nil
	}
	// Empty or unrelated HTTP 200 is not a successful API response.
	if b, ok := v["success"]; ok && string(b) == "true" {
		return VMControlResult{State: "succeeded", JobID: job}, nil
	}
	return VMControlResult{}, ErrAsyncResponse
}
