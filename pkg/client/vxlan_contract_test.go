package client

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCreateL2VxlanNetworkPool_RequestContract(t *testing.T) {
	description := "test pool"
	physicalInterface := "bond0"
	tagUuids := []string{"tag-1", "tag-2"}
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/l2-networks/vxlan-pool" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}

		body := decodeRequestBody(t, r)
		paramsBody, ok := body["params"].(map[string]interface{})
		if !ok {
			t.Fatalf("expected params object, got %#v", body["params"])
		}
		assertJSONField(t, paramsBody, "name", "pool-1")
		assertJSONField(t, paramsBody, "description", description)
		assertJSONField(t, paramsBody, "zoneUuid", "zone-1")
		assertJSONField(t, paramsBody, "physicalInterface", "bond0")
		assertJSONStringArrayField(t, paramsBody, "tagUuid", tagUuids...)
		if _, found := paramsBody["tagUuids"]; found {
			t.Fatalf("params must use the singular tagUuid wire key: %#v", paramsBody)
		}
		assertJSONStringArrayField(t, body, "systemTags", "system-tag")
		if _, found := body["createL2VxlanNetworkPool"]; found {
			t.Fatalf("request must not contain a createL2VxlanNetworkPool wrapper: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"pool-uuid","name":"pool-1","zoneUuid":"zone-1","physicalInterface":"bond0"}}`))
	})

	pool, err := cli.CreateL2VxlanNetworkPool(param.CreateL2VxlanNetworkPoolParam{
		BaseParam: param.BaseParam{SystemTags: []string{"system-tag"}},
		Params: param.CreateL2VxlanNetworkPoolParamDetail{
			Name:              "pool-1",
			Description:       &description,
			ZoneUuid:          "zone-1",
			PhysicalInterface: &physicalInterface,
			TagUuids:          tagUuids,
		},
	})
	if err != nil {
		t.Fatalf("CreateL2VxlanNetworkPool returned error: %v", err)
	}
	if pool.UUID != "pool-uuid" || pool.Name != "pool-1" {
		t.Fatalf("unexpected pool inventory: %+v", pool)
	}
}

func TestAttachL2NetworkToCluster_RequestContract(t *testing.T) {
	providerType := "LinuxBridge"
	vtepCIDRTag := "l2NetworkUuid::pool-1::clusterUuid::cluster-1::cidr::{172.25.0.0/16}"
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/l2-networks/pool-1/clusters/cluster-1" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}

		body := decodeRequestBody(t, r)
		assertJSONField(t, body, "l2ProviderType", providerType)
		assertJSONStringArrayField(t, body, "systemTags", vtepCIDRTag)
		if _, found := body["attachL2NetworkToCluster"]; found {
			t.Fatalf("request must not contain an attachL2NetworkToCluster wrapper: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"pool-1","attachedClusterUuids":["cluster-1"]}}`))
	})

	_, err := cli.AttachL2NetworkToCluster("pool-1", "cluster-1", param.AttachL2NetworkToClusterParam{
		BaseParam: param.BaseParam{SystemTags: []string{vtepCIDRTag}},
		Params: param.AttachL2NetworkToClusterParamDetail{
			L2ProviderType: &providerType,
		},
	})
	if err != nil {
		t.Fatalf("AttachL2NetworkToCluster returned error: %v", err)
	}
}

func TestUpdateVniRange_RequestContract(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/zstack/v1/l2-networks/vxlan-pool/vni-ranges/range-1" {
			t.Fatalf("unexpected request path: %s", r.URL.Path)
		}

		body := decodeRequestBody(t, r)
		updateBody, ok := body["updateVniRange"].(map[string]interface{})
		if !ok {
			t.Fatalf("expected updateVniRange object, got %#v", body["updateVniRange"])
		}
		assertJSONField(t, updateBody, "name", "range-renamed")
		assertJSONStringArrayField(t, body, "systemTags", "system-tag")
		if _, found := body["name"]; found {
			t.Fatalf("request must not contain a top-level name field: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"range-1","name":"range-renamed","startVni":10,"endVni":20}}`))
	})

	rangeInventory, err := cli.UpdateVniRange("range-1", param.UpdateVniRangeParam{
		BaseParam: param.BaseParam{SystemTags: []string{"system-tag"}},
		Params: param.UpdateVniRangeParamDetail{
			Name: "range-renamed",
		},
	})
	if err != nil {
		t.Fatalf("UpdateVniRange returned error: %v", err)
	}
	if rangeInventory.UUID != "range-1" || rangeInventory.Name != "range-renamed" {
		t.Fatalf("unexpected VNI range inventory: %+v", rangeInventory)
	}
}

func decodeRequestBody(t *testing.T, r *http.Request) map[string]interface{} {
	t.Helper()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("read request body: %v", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		t.Fatalf("decode request body %q: %v", bodyBytes, err)
	}
	return body
}

func assertJSONField(t *testing.T, object map[string]interface{}, key string, want interface{}) {
	t.Helper()
	got, found := object[key]
	if !found {
		t.Fatalf("missing JSON field %q in %#v", key, object)
	}
	if got != want {
		t.Fatalf("unexpected JSON field %q: got %#v, want %#v", key, got, want)
	}
}

func assertJSONStringArrayField(t *testing.T, object map[string]interface{}, key string, want ...string) {
	t.Helper()
	raw, found := object[key]
	if !found {
		t.Fatalf("missing JSON field %q in %#v", key, object)
	}
	got, ok := raw.([]interface{})
	if !ok || len(got) != len(want) {
		t.Fatalf("unexpected JSON array field %q: got %#v, want %#v", key, raw, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("unexpected JSON array field %q: got %#v, want %#v", key, raw, want)
		}
	}
}
