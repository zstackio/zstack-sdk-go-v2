// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestAddModelCenterParam_MarshalJSON(t *testing.T) {
	p := param.AddModelCenterParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddModelCenterParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddModelCenterParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateModelCenterParam_MarshalJSON(t *testing.T) {
	p := param.UpdateModelCenterParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateModelCenterParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateModelCenterParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

