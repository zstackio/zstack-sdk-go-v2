// Copyright (c) ZStack.io, Inc.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// Errors never include request parameters, credentials, endpoints or response bodies.
var ErrAsyncResponse = errors.New("invalid asynchronous API response")
var ErrAsyncRequest = errors.New("asynchronous API request failed")
var ErrAsyncNotFound = errors.New("asynchronous API result not found")

type VMAsyncResult struct {
	State     string // accepted, running, succeeded, failed; errors leave the result unknown
	JobID     string
	Inventory *view.VmInstanceInventoryView
}

func asyncID(s string) bool {
	if len(s) != 32 {
		return false
	}
	for _, c := range s {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F') {
			return false
		}
	}
	return true
}

// A location may name an internal management node. Extract only a strict job ID;
// never send credentials to the returned authority or persist its URL.
func (cli *ZSHttpClient) asyncJobID(location string) (string, error) {
	u, err := url.Parse(location)
	prefix := "/" + strings.Trim(cli.contextPath, "/") + "/v1/api-jobs/"
	if cli.contextPath == "" {
		prefix = "/v1/api-jobs/"
	}
	if err != nil || u.User != nil || u.RawQuery != "" || u.Fragment != "" || u.RawPath != "" || strings.Contains(location, "?") || strings.Contains(location, "#") || (u.Scheme != "" && u.Scheme != "http" && u.Scheme != "https") || (u.Scheme == "" && u.Host != "") || !strings.HasPrefix(u.Path, prefix) {
		return "", ErrAsyncResponse
	}
	id := strings.TrimPrefix(u.Path, prefix)
	if !asyncID(id) {
		return "", ErrAsyncResponse
	}
	return strings.ToLower(id), nil
}

// CreateVmInstanceAsync sends exactly one POST and never polls or retries.
// A failed call does not prove that the Cloud did not accept the request.
func (cli *ZSClient) CreateVmInstanceAsync(ctx context.Context, p param.CreateVmInstanceParam) (VMAsyncResult, error) {
	if cli.readOnly {
		return VMAsyncResult{}, ErrAsyncRequest
	}
	b, err := json.Marshal(p)
	if err != nil {
		return VMAsyncResult{}, ErrAsyncRequest
	}
	return cli.vmAsyncRequest(ctx, http.MethodPost, cli.getRequestURL("v1/vm-instances"), b, "")
}

// QueryVmCreation performs one GET, independent of configured retry counts.
// 404 means the result is unavailable, not that the operation failed.
func (cli *ZSClient) QueryVmCreation(ctx context.Context, jobID string) (VMAsyncResult, error) {
	if !asyncID(jobID) {
		return VMAsyncResult{}, ErrAsyncRequest
	}
	return cli.vmAsyncRequest(ctx, http.MethodGet, cli.getRequestURL("v1/api-jobs/"+strings.ToLower(jobID)), nil, strings.ToLower(jobID))
}

func (cli *ZSHttpClient) vmAsyncRequest(ctx context.Context, method, target string, body []byte, job string) (VMAsyncResult, error) {
	fail := func(e error) (VMAsyncResult, error) { return VMAsyncResult{}, e }
	header, err := cli.getHeader(target, method)
	if err != nil {
		return fail(ErrAsyncRequest)
	}
	req, err := http.NewRequestWithContext(ctx, method, target, bytes.NewReader(body))
	if err != nil {
		return fail(ErrAsyncRequest)
	}
	req.Header = header
	req.Header.Set("Content-Type", "application/json")
	httpClient := *cli.httpClient
	httpClient.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	resp, err := httpClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return fail(ctx.Err())
		}
		return fail(ErrAsyncRequest)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound && method == http.MethodGet {
		return fail(ErrAsyncNotFound)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 && resp.StatusCode != 503 {
		return fail(ErrAsyncResponse)
	}
	b, err := io.ReadAll(io.LimitReader(resp.Body, 2*1024*1024+1))
	if err != nil || len(b) > 2*1024*1024 {
		return fail(ErrAsyncResponse)
	}
	// A running job legitimately returns an empty 202 body.
	if method == http.MethodGet && resp.StatusCode == 202 && len(bytes.TrimSpace(b)) == 0 {
		return VMAsyncResult{State: "running", JobID: job}, nil
	}
	var data map[string]json.RawMessage
	if json.Unmarshal(b, &data) != nil || data == nil {
		return fail(ErrAsyncResponse)
	}
	if raw, ok := data["error"]; ok {
		var cloudError struct {
			Code string `json:"code"`
		}
		if (resp.StatusCode != 200 && resp.StatusCode != 503) || json.Unmarshal(raw, &cloudError) != nil || cloudError.Code == "" {
			return fail(ErrAsyncResponse)
		}
		return VMAsyncResult{State: "failed", JobID: job}, nil
	}
	if resp.StatusCode == 503 {
		return fail(ErrAsyncResponse)
	}
	if raw, ok := data["success"]; ok && string(raw) != "true" {
		return fail(ErrAsyncResponse)
	}
	if resp.StatusCode == 202 {
		var location string
		if raw, ok := data["location"]; ok {
			if json.Unmarshal(raw, &location) != nil {
				return fail(ErrAsyncResponse)
			}
			id, e := cli.asyncJobID(location)
			if e != nil {
				return fail(e)
			}
			if job != "" && job != id {
				return fail(ErrAsyncResponse)
			}
			job = id
		} else if method == http.MethodPost {
			return fail(ErrAsyncResponse)
		}
		state := "running"
		if method == http.MethodPost {
			state = "accepted"
		}
		return VMAsyncResult{State: state, JobID: job}, nil
	}
	if _, ok := data["location"]; ok {
		return fail(ErrAsyncResponse)
	}
	var vm view.VmInstanceInventoryView
	if json.Unmarshal(data["inventory"], &vm) != nil || !asyncID(vm.UUID) {
		return fail(ErrAsyncResponse)
	}
	return VMAsyncResult{State: "succeeded", JobID: job, Inventory: &vm}, nil
}
