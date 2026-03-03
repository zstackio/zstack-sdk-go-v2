// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateSNSUniversalSmsEndpointParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSNSUniversalSmsEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSNSUniversalSmsEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSNSUniversalSmsEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateSNSUniversalSmsEndpointParam_MarshalJSON(t *testing.T) {
	p := param.CreateSNSUniversalSmsEndpointParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSNSUniversalSmsEndpointParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSNSUniversalSmsEndpointParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

