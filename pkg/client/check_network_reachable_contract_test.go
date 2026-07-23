package client

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCheckNetworkReachable_RequestContract(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/zops/check/network" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		if len(body) != 0 {
			t.Fatalf("GET request must not contain a body, got %q", body)
		}

		query := r.URL.Query()
		assertStringValues(t, query["sourceHostnames"], "172.25.16.124", "172.25.16.125")
		assertStringValues(t, query["targetHostnames"], "ntp.example.com", "8.8.8.8")
		assertStringValues(t, query["systemTags"], "system-tag-1", "system-tag-2")
		assertStringValues(t, query["userTags"], "user-tag")
		if got := query.Get("requestIp"); got != "172.25.16.10" {
			t.Fatalf("unexpected requestIp: got %q", got)
		}
		if strings.Contains(r.URL.RawQuery, "%5B") || strings.Contains(r.URL.RawQuery, "%5D") {
			t.Fatalf("list parameters must use repeated query keys, got %s", r.URL.RawQuery)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"results":[{"sourceHostname":"172.25.16.124","targetHostname":"ntp.example.com","status":"Connected"}],"success":true}`))
	})

	result, err := cli.CheckNetworkReachable(param.CheckNetworkReachableParam{
		BaseParam: param.BaseParam{
			SystemTags: []string{"system-tag-1", "system-tag-2"},
			UserTags:   []string{"user-tag"},
			RequestIp:  "172.25.16.10",
		},
		Params: param.CheckNetworkReachableParamDetail{
			SourceHostnames: []string{"172.25.16.124", "172.25.16.125"},
			TargetHostnames: []string{"ntp.example.com", "8.8.8.8"},
		},
	})
	if err != nil {
		t.Fatalf("CheckNetworkReachable returned error: %v", err)
	}
	if !result.Success || len(result.Results) != 1 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Results[0].SourceHostname != "172.25.16.124" ||
		result.Results[0].TargetHostname != "ntp.example.com" ||
		result.Results[0].Status != "Connected" {
		t.Fatalf("unexpected result item: %+v", result.Results[0])
	}
}

func TestCheckNetworkReachable_SourceHostnamesOptional(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if _, found := query["sourceHostnames"]; found {
			t.Fatalf("sourceHostnames must be omitted when unset: %s", r.URL.RawQuery)
		}
		assertStringValues(t, query["targetHostnames"], "ntp.example.com")

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"results":[],"success":true}`))
	})

	_, err := cli.CheckNetworkReachable(param.CheckNetworkReachableParam{
		Params: param.CheckNetworkReachableParamDetail{
			TargetHostnames: []string{"ntp.example.com"},
		},
	})
	if err != nil {
		t.Fatalf("CheckNetworkReachable returned error: %v", err)
	}
}

func TestCheckNetworkReachable_TargetHostnamesRequired(t *testing.T) {
	requestSent := false
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		requestSent = true
		t.Fatal("request must not be sent without targetHostnames")
	})

	_, err := cli.CheckNetworkReachable(param.CheckNetworkReachableParam{})
	if err == nil {
		t.Fatal("expected missing targetHostnames error")
	}
	if !strings.Contains(err.Error(), "targetHostnames is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if requestSent {
		t.Fatal("request was sent without targetHostnames")
	}
}

func assertStringValues(t *testing.T, got []string, want ...string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected query values: got %#v, want %#v", got, want)
	}
}
