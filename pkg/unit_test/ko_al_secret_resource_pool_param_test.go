// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateKoAlSecretResourcePoolParam_MarshalJSON(t *testing.T) {
	p := param.UpdateKoAlSecretResourcePoolParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateKoAlSecretResourcePoolParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateKoAlSecretResourcePoolParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateKoAlSecretResourcePoolParam_MarshalJSON(t *testing.T) {
	p := param.CreateKoAlSecretResourcePoolParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateKoAlSecretResourcePoolParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateKoAlSecretResourcePoolParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

