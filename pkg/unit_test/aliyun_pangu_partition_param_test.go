// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestAddAliyunPanguPartitionParam_MarshalJSON(t *testing.T) {
	p := param.AddAliyunPanguPartitionParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddAliyunPanguPartitionParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddAliyunPanguPartitionParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateAliyunPanguPartitionParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAliyunPanguPartitionParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAliyunPanguPartitionParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAliyunPanguPartitionParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

