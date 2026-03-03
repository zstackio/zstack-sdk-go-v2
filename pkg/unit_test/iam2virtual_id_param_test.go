// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCreateIAM2VirtualIDParam_MarshalJSON(t *testing.T) {
	p := param.CreateIAM2VirtualIDParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateIAM2VirtualIDParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateIAM2VirtualIDParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateIAM2VirtualIDParam_MarshalJSON(t *testing.T) {
	p := param.UpdateIAM2VirtualIDParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateIAM2VirtualIDParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateIAM2VirtualIDParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

