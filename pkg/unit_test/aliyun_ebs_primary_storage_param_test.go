// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateAliyunEbsPrimaryStorageParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAliyunEbsPrimaryStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAliyunEbsPrimaryStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAliyunEbsPrimaryStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddAliyunEbsPrimaryStorageParam_MarshalJSON(t *testing.T) {
	p := param.AddAliyunEbsPrimaryStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddAliyunEbsPrimaryStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddAliyunEbsPrimaryStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

