// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateSAML2ClientParam_MarshalJSON(t *testing.T) {
	p := param.UpdateSAML2ClientParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateSAML2ClientParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateSAML2ClientParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestCreateSAML2ClientParam_MarshalJSON(t *testing.T) {
	p := param.CreateSAML2ClientParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateSAML2ClientParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateSAML2ClientParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

