// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCreateGuestVmScriptParam_MarshalJSON(t *testing.T) {
	p := param.CreateGuestVmScriptParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateGuestVmScriptParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateGuestVmScriptParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateGuestVmScriptParam_MarshalJSON(t *testing.T) {
	p := param.UpdateGuestVmScriptParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateGuestVmScriptParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateGuestVmScriptParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

