// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestAddAppBuildSystemParam_MarshalJSON(t *testing.T) {
	p := param.AddAppBuildSystemParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddAppBuildSystemParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddAppBuildSystemParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateAppBuildSystemParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAppBuildSystemParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAppBuildSystemParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAppBuildSystemParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

