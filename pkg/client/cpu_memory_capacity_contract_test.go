package client

import (
	"net/http"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestGetCpuMemoryCapacityWithParams_All(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/zstack/v1/hosts/capacities/cpu-memory" || r.URL.RawQuery != "all=true" {
			t.Fatalf("unexpected capacity request: %s %s", r.Method, r.URL.RequestURI())
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"totalCpu":4,"availableCpu":2,"totalMemory":8,"availableMemory":4}`))
	})

	capacity, err := cli.GetCpuMemoryCapacityWithParams(param.GetCpuMemoryCapacityParamDetail{All: true})
	if err != nil {
		t.Fatalf("GetCpuMemoryCapacityWithParams: %v", err)
	}
	if capacity.TotalCpu != 4 || capacity.AvailableCpu != 2 || capacity.TotalMemory != 8 || capacity.AvailableMemory != 4 {
		t.Fatalf("unexpected capacity: %+v", capacity)
	}
}

func TestGetCpuMemoryCapacityWithParams_ZoneScope(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/zstack/v1/hosts/capacities/cpu-memory" {
			t.Fatalf("unexpected capacity request: %s %s", r.Method, r.URL.RequestURI())
		}
		query := r.URL.Query()
		assertStringValues(t, query["zoneUuids"], "zone-1", "zone-2")
		if _, found := query["all"]; found {
			t.Fatalf("all must be omitted for a zone-scoped query: %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"totalCpu":4,"availableCpu":2,"totalMemory":8,"availableMemory":4}`))
	})

	_, err := cli.GetCpuMemoryCapacityWithParams(param.GetCpuMemoryCapacityParamDetail{ZoneUuids: []string{"zone-1", "zone-2"}})
	if err != nil {
		t.Fatalf("GetCpuMemoryCapacityWithParams: %v", err)
	}
}
