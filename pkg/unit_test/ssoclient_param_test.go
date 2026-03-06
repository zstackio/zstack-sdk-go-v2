// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestDeleteSSOClientParam_MarshalJSON(t *testing.T) {
	p := param.DeleteSSOClientParam{
		Params: param.DeleteSSOClientParamDetail{
			Uuid: "test-uuid-001",
		},
	}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
	// Verify params.uuid is present
	params, ok := raw["params"].(map[string]interface{})
	if !ok {
		t.Fatal("expected 'params' key in marshaled JSON")
	}
	if params["uuid"] != "test-uuid-001" {
		t.Errorf("expected uuid 'test-uuid-001', got %v", params["uuid"])
	}
}

func TestDeleteSSOClientParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{"params":{"uuid":"test-uuid-001"}}`
	var p param.DeleteSSOClientParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
	if p.Params.Uuid != "test-uuid-001" {
		t.Errorf("expected uuid 'test-uuid-001', got %v", p.Params.Uuid)
	}
}

func TestGetSSOClientParam_MarshalJSON(t *testing.T) {
	p := param.GetSSOClientParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestGetSSOClientParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.GetSSOClientParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}
