// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateContainerManagementEndpointParam_MarshalJSON(t *testing.T) {
	p := param.UpdateContainerManagementEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateContainerManagementEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateContainerManagementEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddContainerManagementEndpointParam_MarshalJSON(t *testing.T) {
	p := param.AddContainerManagementEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddContainerManagementEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddContainerManagementEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

