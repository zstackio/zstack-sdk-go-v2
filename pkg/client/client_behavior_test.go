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
