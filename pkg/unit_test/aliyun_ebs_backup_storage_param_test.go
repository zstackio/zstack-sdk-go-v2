// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestAddAliyunEbsBackupStorageParam_MarshalJSON(t *testing.T) {
	p := param.AddAliyunEbsBackupStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddAliyunEbsBackupStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddAliyunEbsBackupStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateAliyunEbsBackupStorageParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAliyunEbsBackupStorageParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAliyunEbsBackupStorageParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAliyunEbsBackupStorageParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

