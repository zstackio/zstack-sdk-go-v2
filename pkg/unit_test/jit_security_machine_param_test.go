// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateJitSecurityMachineParam_MarshalJSON(t *testing.T) {
	p := param.UpdateJitSecurityMachineParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateJitSecurityMachineParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateJitSecurityMachineParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddJitSecurityMachineParam_MarshalJSON(t *testing.T) {
	p := param.AddJitSecurityMachineParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddJitSecurityMachineParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddJitSecurityMachineParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

