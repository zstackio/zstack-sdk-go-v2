// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCreatePluginSecretResourcePoolParam_MarshalJSON(t *testing.T) {
	p := param.CreatePluginSecretResourcePoolParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreatePluginSecretResourcePoolParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreatePluginSecretResourcePoolParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdatePluginSecretResourcePoolParam_MarshalJSON(t *testing.T) {
	p := param.UpdatePluginSecretResourcePoolParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdatePluginSecretResourcePoolParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdatePluginSecretResourcePoolParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

