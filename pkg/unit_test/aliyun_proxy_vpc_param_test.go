// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestCreateAliyunProxyVpcParam_MarshalJSON(t *testing.T) {
	p := param.CreateAliyunProxyVpcParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestCreateAliyunProxyVpcParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.CreateAliyunProxyVpcParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestUpdateAliyunProxyVpcParam_MarshalJSON(t *testing.T) {
	p := param.UpdateAliyunProxyVpcParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateAliyunProxyVpcParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateAliyunProxyVpcParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

