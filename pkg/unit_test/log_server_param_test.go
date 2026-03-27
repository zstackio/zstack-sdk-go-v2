// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateLogServerParam_MarshalJSON(t *testing.T) {
	p := param.UpdateLogServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateLogServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateLogServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddLogServerParam_MarshalJSON(t *testing.T) {
	p := param.AddLogServerParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddLogServerParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddLogServerParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

