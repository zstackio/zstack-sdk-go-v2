// Copyright (c) ZStack.io, Inc.
// Auto-generated view tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

func TestAliyunNasFileSystemInventoryView_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"uuid": "test-uuid-001",
		"name": "test-aliyun_nas_file_system",
		"createDate": "2024-01-01T00:00:00.000+08:00",
		"lastOpDate": "2024-01-01T00:00:00.000+08:00"
	}`
	var v view.AliyunNasFileSystemInventoryView
	err := json.Unmarshal([]byte(jsonStr), &v)
	assertNoError(t, err)
}

func TestAliyunNasFileSystemInventoryView_UnmarshalEmpty(t *testing.T) {
	var v view.AliyunNasFileSystemInventoryView
	err := json.Unmarshal([]byte(`{}`), &v)
	assertNoError(t, err)
}

