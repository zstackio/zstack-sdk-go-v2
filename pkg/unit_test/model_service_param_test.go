// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateModelServiceParam_MarshalJSON(t *testing.T) {
	p := param.UpdateModelServiceParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateModelServiceParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateModelServiceParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddModelServiceParam_MarshalJSON(t *testing.T) {
	p := param.AddModelServiceParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddModelServiceParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddModelServiceParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

