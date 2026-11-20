// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestDetachTagFromResources_RequestContract(t *testing.T) {
	for _, mode := range []param.DeleteMode{param.DeleteModePermissive, param.DeleteModeEnforcing, ""} {
		t.Run(string(mode), func(t *testing.T) {
			resourceUuids := []string{"resource-a", "resource-b"}
			requests := 0
			cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
				requests++
				if r.Method != http.MethodDelete || r.URL.Path != "/zstack/v1/tags/tag-1/resources" {
					t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
				}
				want := url.Values{"resourceUuids": resourceUuids}
				if mode != "" {
					want.Set("deleteMode", string(mode))
				}
				if got := r.URL.Query(); !reflect.DeepEqual(got, want) {
					t.Fatalf("unexpected query: got %#v, want %#v", got, want)
				}
				body, err := io.ReadAll(r.Body)
				if err != nil || len(body) != 0 {
					t.Fatalf("expected empty DELETE body, got %q, error %v", body, err)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{}`))
			})
			if err := cli.DetachTagFromResources("tag-1", mode, resourceUuids); err != nil {
				t.Fatalf("DetachTagFromResources: %v", err)
			}
			if requests != 1 {
				t.Fatalf("expected one request, got %d", requests)
			}
		})
	}
}

func TestDetachTagFromResources_RejectsMissingTargets(t *testing.T) {
	for _, tc := range []struct {
		name          string
		tagUuid       string
		resourceUuids []string
	}{
		{"missing tag", "", []string{"resource-a"}},
		{"nil resources", "tag-1", nil},
		{"empty resources", "tag-1", []string{}},
		{"empty resource", "tag-1", []string{"resource-a", ""}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// A nil client ensures validation cannot accidentally issue a request.
			var cli *ZSClient
			if err := cli.DetachTagFromResources(tc.tagUuid, param.DeleteModePermissive, tc.resourceUuids); err == nil {
				t.Fatal("expected missing target error")
			}
		})
	}
}

func TestDetachTagFromResources_WaitsForAsyncCompletion(t *testing.T) {
	deletes, polls := 0, 0
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodDelete && r.URL.Path == "/zstack/v1/tags/tag-1/resources":
			deletes++
			w.WriteHeader(http.StatusAccepted)
			_, _ = fmt.Fprintf(w, `{"location":"http://%s/zstack/v1/api-jobs/job-1"}`, r.Host)
		case r.Method == http.MethodGet && r.URL.Path == "/zstack/v1/api-jobs/job-1":
			polls++
			_, _ = w.Write([]byte(`{}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
		}
	})
	if err := cli.DetachTagFromResources("tag-1", param.DeleteModePermissive, []string{"resource-a", "resource-b"}); err != nil {
		t.Fatalf("DetachTagFromResources: %v", err)
	}
	if deletes != 1 || polls != 1 {
		t.Fatalf("expected one DELETE and one completion poll, got %d and %d", deletes, polls)
	}
}

func TestDetachTagFromResources_PropagatesAPIError(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"code":"TEST.DETACH","description":"detach denied","details":"permission denied"}}`))
	})
	err := cli.DetachTagFromResources("tag-1", param.DeleteModePermissive, []string{"resource-a"})
	if err == nil || !strings.Contains(err.Error(), "TEST.DETACH") {
		t.Fatalf("expected API error to reach caller, got %v", err)
	}
}
