// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCreateDatasetParam_MarshalJSON(t *testing.T) {
	p := param.CreateDatasetParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateDatasetParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateDatasetParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateDatasetParam_MarshalJSON(t *testing.T) {
	p := param.UpdateDatasetParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateDatasetParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateDatasetParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

