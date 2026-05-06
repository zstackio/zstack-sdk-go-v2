package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/util/jsonutils"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

func newTestZSClient(t *testing.T, handler http.HandlerFunc) *ZSClient {
	t.Helper()

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}

	port, err := strconv.Atoi(serverURL.Port())
	if err != nil {
		t.Fatalf("parse server port: %v", err)
	}

	config := NewZSConfig(serverURL.Hostname(), port, "zstack").AccessKey("ak", "sk")
	return NewZSClient(config)
}

func TestPutWithRespKey_EmptyResponseKeyFallsBackToInventoryEnvelope(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/iam2/projects/project-1/actions" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"project-1","name":"project-name"}}`))
	})

	var result view.IAM2ProjectInventoryView

	err := cli.PutWithRespKey("v1/iam2/projects", "project-1", "", map[string]interface{}{
		"updateIAM2Project": map[string]interface{}{"name": "project-name"},
	}, &result)
	if err != nil {
		t.Fatalf("PutWithRespKey returned error: %v", err)
	}

	if result.UUID != "project-1" || result.Name != "project-name" {
		t.Fatalf("expected inventory fields to be unmarshaled, got %+v", result)
	}
}

func TestStandardJSONUnmarshal_InventoryIntoIAM2ProjectView(t *testing.T) {
	var result view.IAM2ProjectInventoryView
	err := json.Unmarshal([]byte(`{"uuid":"project-1","name":"project-name"}`), &result)
	if err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}
	if result.UUID != "project-1" || result.Name != "project-name" {
		t.Fatalf("expected standard json unmarshal to populate result, got %+v", result)
	}
}

func TestJSONUtilsInventoryString_IsStandardJSON(t *testing.T) {
	body, err := jsonutils.Parse([]byte(`{"inventory":{"uuid":"project-1","name":"project-name"}}`))
	if err != nil {
		t.Fatalf("jsonutils.Parse returned error: %v", err)
	}
	inventory, err := body.Get("inventory")
	if err != nil {
		t.Fatalf("Get inventory returned error: %v", err)
	}

	var result view.IAM2ProjectInventoryView
	err = json.Unmarshal([]byte(inventory.String()), &result)
	if err != nil {
		t.Fatalf("json.Unmarshal inventory string returned error: %v", err)
	}
	if result.UUID != "project-1" || result.Name != "project-name" {
		t.Fatalf("expected inventory string to decode into result, got %+v from %s", result, inventory.String())
	}
}

func TestHttpPut_ResponseContainsInventoryEnvelope(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"project-1","name":"project-name"}}`))
	})

	_, _, resp, err := cli.httpPut(t.Context(), cli.getPutURL("v1/iam2/projects", "project-1", "actions"), jsonMarshal(map[string]interface{}{
		"updateIAM2Project": map[string]interface{}{"name": "project-name"},
	}), false)
	if err != nil {
		t.Fatalf("httpPut returned error: %v", err)
	}
	if !resp.Contains(responseKeyInventory) {
		t.Fatalf("expected response to contain inventory, got %s", resp.String())
	}

	inventory, err := resp.Get(responseKeyInventory)
	if err != nil {
		t.Fatalf("Get inventory returned error: %v", err)
	}

	var result view.IAM2ProjectInventoryView
	err = json.Unmarshal([]byte(inventory.String()), &result)
	if err != nil {
		t.Fatalf("json.Unmarshal inventory returned error: %v", err)
	}
	if result.UUID != "project-1" || result.Name != "project-name" {
		t.Fatalf("expected decoded inventory, got %+v from %s", result, inventory.String())
	}
}

func TestDeleteAndExpungeIAM2Project_DeletesThenExpunges(t *testing.T) {
	requests := make([]string, 0, 2)
	bodies := make([]string, 0, 2)

	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}

		requests = append(requests, r.Method+" "+r.URL.RequestURI())
		bodies = append(bodies, string(body))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	})

	err := cli.DeleteAndExpungeIAM2Project("project-1", param.DeleteModePermissive)
	if err != nil {
		t.Fatalf("DeleteAndExpungeIAM2Project returned error: %v", err)
	}

	if len(requests) != 2 {
		t.Fatalf("expected 2 requests, got %d (%v)", len(requests), requests)
	}

	if requests[0] != "DELETE /zstack/v1/iam2/projects/project-1?deleteMode=Permissive" {
		t.Fatalf("unexpected delete request: %s", requests[0])
	}

	if requests[1] != "PUT /zstack/v1/iam2/projects/project-1/actions" {
		t.Fatalf("unexpected expunge request: %s", requests[1])
	}

	if !strings.Contains(bodies[1], `"expungeIAM2Project":{}`) {
		t.Fatalf("expected expunge body, got %s", bodies[1])
	}
}

func TestDeleteDirectory_UsesActionURLAndBody(t *testing.T) {
	var request string
	var body string

	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}

		request = r.Method + " " + r.URL.RequestURI()
		body = string(bodyBytes)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	})

	err := cli.DeleteDirectory("directory-1", param.DeleteModePermissive)
	if err != nil {
		t.Fatalf("DeleteDirectory returned error: %v", err)
	}

	if request != "DELETE /zstack/v1/delete/directory" {
		t.Fatalf("unexpected delete request: %s", request)
	}

	var payload struct {
		DeleteDirectory struct {
			UUID       string `json:"uuid"`
			DeleteMode string `json:"deleteMode"`
		} `json:"deleteDirectory"`
	}
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		t.Fatalf("unmarshal deleteDirectory body: %v", err)
	}
	if payload.DeleteDirectory.UUID != "directory-1" || payload.DeleteDirectory.DeleteMode != "Permissive" {
		t.Fatalf("expected deleteDirectory body, got %s", body)
	}
}

// TestPutWithRespKey_EmptyResponseKey_InventoryWithZStackTimeFormat covers
// SDK-BUG-001's regression target: when an Update* action passes an empty
// responseKey but the API returns the inventory envelope, the SDK must
// decode through jsonutils.Unmarshal (which understands timeutils formats
// such as ZStackTimeFormat — "Jan 2, 2006 3:04:05 PM"), not the stdlib
// json.Unmarshal which silently leaves time.Time fields zero-valued.
func TestPutWithRespKey_EmptyResponseKey_InventoryWithZStackTimeFormat(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// lastOpDate is in ZStackTimeFormat — stdlib json.Unmarshal cannot
		// parse this into time.Time but jsonutils.Unmarshal can.
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"project-1","name":"project-name","lastOpDate":"Apr 26, 2026 10:30:00 AM"}}`))
	})

	var result view.IAM2ProjectInventoryView
	err := cli.PutWithRespKey("v1/iam2/projects", "project-1", "", map[string]interface{}{
		"updateIAM2Project": map[string]interface{}{"name": "project-name"},
	}, &result)
	if err != nil {
		t.Fatalf("PutWithRespKey returned error: %v", err)
	}

	if result.UUID != "project-1" || result.Name != "project-name" {
		t.Fatalf("expected inventory fields to be unmarshaled, got %+v", result)
	}
	if result.LastOpDate.IsZero() {
		t.Fatalf("expected LastOpDate to be parsed from ZStackTimeFormat, got zero value")
	}
	// Sanity-check the parsed components rather than an exact timezone-bound
	// equality so the test stays portable.
	if result.LastOpDate.Year() != 2026 || result.LastOpDate.Month() != time.April || result.LastOpDate.Day() != 26 {
		t.Fatalf("expected 2026-04-26, got %v", result.LastOpDate)
	}
}

// TestPutWithRespKey_EmptyResponseKey_NoInventoryEnvelopeFallsThrough verifies
// the bounded-risk guarantee: when the response does NOT include an inventory
// envelope, the SDK keeps its original whole-body Unmarshal path. This protects
// any non-Update PutWithRespKey(..., "", ...) callers whose responses are bare
// objects, "results"-shaped, scalar success blobs, etc.
func TestPutWithRespKey_EmptyResponseKey_NoInventoryEnvelopeFallsThrough(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// No `inventory` key — bare object the SDK must continue to decode
		// directly into the target struct.
		_, _ = w.Write([]byte(`{"uuid":"bare-1","name":"bare-name"}`))
	})

	var result view.IAM2ProjectInventoryView
	err := cli.PutWithRespKey("v1/iam2/projects", "bare-1", "", map[string]interface{}{
		"updateIAM2Project": map[string]interface{}{"name": "bare-name"},
	}, &result)
	if err != nil {
		t.Fatalf("PutWithRespKey returned error: %v", err)
	}

	if result.UUID != "bare-1" || result.Name != "bare-name" {
		t.Fatalf("expected fallback whole-body unmarshal to populate result, got %+v", result)
	}
}

// TestPutWithRespKey_ExplicitResponseKey_StillWorks asserts the fix did not
// alter the path used by callers that already pass an explicit responseKey
// (e.g. anything that goes through Put or PutWithRespKey with "inventory").
func TestPutWithRespKey_ExplicitResponseKey_StillWorks(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"project-1","name":"explicit"}}`))
	})

	var result view.IAM2ProjectInventoryView
	err := cli.PutWithRespKey("v1/iam2/projects", "project-1", responseKeyInventory, map[string]interface{}{
		"updateIAM2Project": map[string]interface{}{"name": "explicit"},
	}, &result)
	if err != nil {
		t.Fatalf("PutWithRespKey returned error: %v", err)
	}
	if result.UUID != "project-1" || result.Name != "explicit" {
		t.Fatalf("expected explicit responseKey to keep working, got %+v", result)
	}
}

// TestPostWithRespKey_EmptyResponseKey_InventoryWithZStackTimeFormat keeps
// PostWithAsync's fallback in lockstep with PutWithAsync so the same bounded
// fix applies to any future Create-path caller that mirrors the Update pattern.
func TestPostWithRespKey_EmptyResponseKey_InventoryWithZStackTimeFormat(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST request, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"project-1","name":"created","lastOpDate":"Apr 26, 2026 10:30:00 AM"}}`))
	})

	var result view.IAM2ProjectInventoryView
	err := cli.PostWithRespKey(t.Context(), "v1/iam2/projects", "", map[string]interface{}{
		"createIAM2Project": map[string]interface{}{"name": "created"},
	}, &result)
	if err != nil {
		t.Fatalf("PostWithRespKey returned error: %v", err)
	}
	if result.UUID != "project-1" || result.Name != "created" {
		t.Fatalf("expected inventory fields to be unmarshaled, got %+v", result)
	}
	if result.LastOpDate.IsZero() {
		t.Fatalf("expected LastOpDate to be parsed from ZStackTimeFormat, got zero value")
	}
}

// TestPostWithRespKey_EmptyResponseKey_NoInventoryEnvelopeFallsThrough mirrors
// the Put-side fallback test for PostWithAsync.
func TestPostWithRespKey_EmptyResponseKey_NoInventoryEnvelopeFallsThrough(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"uuid":"bare-1","name":"bare-name"}`))
	})

	var result view.IAM2ProjectInventoryView
	err := cli.PostWithRespKey(t.Context(), "v1/iam2/projects", "", map[string]interface{}{
		"createIAM2Project": map[string]interface{}{"name": "bare-name"},
	}, &result)
	if err != nil {
		t.Fatalf("PostWithRespKey returned error: %v", err)
	}
	if result.UUID != "bare-1" || result.Name != "bare-name" {
		t.Fatalf("expected fallback whole-body unmarshal to populate result, got %+v", result)
	}
}

func TestAttachTagToResources_DecodesTopLevelEventResponse(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/tags/tag-1/resources" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		if !strings.Contains(string(body), `"resourceUuids":["res-1"]`) {
			t.Fatalf("expected resourceUuids in request body, got %s", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"success":true,"results":[{"success":true,"inventory":{"uuid":"binding-1","resourceUuid":"res-1"}}]}`))
	})

	resp, err := cli.AttachTagToResources("tag-1", param.AttachTagToResourcesParam{
		Params: param.AttachTagToResourcesParamDetail{
			ResourceUuids: []string{"res-1"},
		},
	})
	if err != nil {
		t.Fatalf("AttachTagToResources returned error: %v", err)
	}
	if !resp.Success || len(resp.Results) != 1 || !resp.Results[0].Success {
		t.Fatalf("expected successful top-level event response, got %+v", resp)
	}
	if resp.Results[0].Inventory.UUID != "binding-1" || resp.Results[0].Inventory.ResourceUuid != "res-1" {
		t.Fatalf("expected decoded result inventory, got %+v", resp.Results[0].Inventory)
	}
}

func TestAddVmNicToSecurityGroup_DecodesTopLevelEventResponse(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/security-groups/sg-1/vm-instances/nics" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		if !strings.Contains(string(body), `"vmNicUuids":["nic-1"]`) {
			t.Fatalf("expected vmNicUuids in request body, got %s", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"success":true}`))
	})

	resp, err := cli.AddVmNicToSecurityGroup("sg-1", param.AddVmNicToSecurityGroupParam{
		Params: param.AddVmNicToSecurityGroupParamDetail{
			VmNicUuids: []string{"nic-1"},
		},
	})
	if err != nil {
		t.Fatalf("AddVmNicToSecurityGroup returned error: %v", err)
	}
	if !resp.Success {
		t.Fatalf("expected successful top-level event response, got %+v", resp)
	}
}
