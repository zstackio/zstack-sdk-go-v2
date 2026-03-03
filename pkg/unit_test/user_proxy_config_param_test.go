// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateUserProxyConfigParam_MarshalJSON(t *testing.T) {
	p := param.UpdateUserProxyConfigParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateUserProxyConfigParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateUserProxyConfigParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateUserProxyConfigParam_MarshalJSON(t *testing.T) {
	p := param.CreateUserProxyConfigParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateUserProxyConfigParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateUserProxyConfigParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

