package client

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

func TestGetEffectiveResourceConfigPathAndValue(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/zstack/v1/resource-configurations/test-resource/kvm/reservedMemory" {
			t.Errorf("wrong resource config request: %s %s", r.Method, r.URL.Path)
		}
		fmt.Fprint(w, `{"value":"2G","effectiveConfigs":[]}`)
	})
	got, err := cli.GetEffectiveResourceConfig("test-resource", "kvm", "reservedMemory")
	if err != nil || got == nil || got.Value != "2G" {
		t.Fatalf("wrong effective config: %+v %v", got, err)
	}
	if len(got.EffectiveConfigs) != 0 {
		t.Fatalf("expected no effective configs, got %+v", got.EffectiveConfigs)
	}
}

func TestGetEffectiveResourceConfigEffectiveConfigs(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"value":"2G","success":true,"effectiveConfigs":[
			{"uuid":"config-host","resourceUuid":"test-resource","resourceType":"HostVO","category":"kvm","name":"reservedMemory","description":"host override","value":"2G"},
			{"uuid":"config-cluster","resourceUuid":"test-cluster","resourceType":"ClusterVO","category":"kvm","name":"reservedMemory","description":"cluster default","value":"1G"}
		]}`)
	})
	got, err := cli.GetEffectiveResourceConfig("test-resource", "kvm", "reservedMemory")
	if err != nil || got == nil {
		t.Fatalf("GetEffectiveResourceConfig: result=%+v error=%v", got, err)
	}
	want := []view.ResourceConfigInventoryView{
		{BaseInfoView: view.BaseInfoView{UUID: "config-host", Name: "reservedMemory"}, ResourceUuid: "test-resource", ResourceType: "HostVO", Category: "kvm", Description: "host override", Value: "2G"},
		{BaseInfoView: view.BaseInfoView{UUID: "config-cluster", Name: "reservedMemory"}, ResourceUuid: "test-cluster", ResourceType: "ClusterVO", Category: "kvm", Description: "cluster default", Value: "1G"},
	}
	if got.Value != "2G" || !got.Success || !reflect.DeepEqual(got.EffectiveConfigs, want) {
		t.Fatalf("unexpected response: %+v; expected value=2G, success=true, effectiveConfigs=%+v", got, want)
	}
}

func TestGetEffectiveResourceConfigErrors(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   string
	}{
		{"not found", http.StatusNotFound, `{"error":{"code":"RESOURCE_NOT_FOUND","details":"resource does not exist"}}`},
		{"server error", http.StatusInternalServerError, `{"error":{"code":"INTERNAL_ERROR","details":"configuration lookup failed"}}`},
		{"invalid JSON", http.StatusOK, `{"value":`},
		{"invalid effectiveConfigs", http.StatusOK, `{"value":"2G","effectiveConfigs":"invalid"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				fmt.Fprint(w, tc.body)
			})
			got, err := cli.GetEffectiveResourceConfig("test-resource", "kvm", "reservedMemory")
			if err == nil || got != nil {
				t.Fatalf("expected nil result and an error, got result=%+v error=%v", got, err)
			}
		})
	}
}
