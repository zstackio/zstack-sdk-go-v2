// Copyright (c) ZStack.io, Inc.
// Auto-generated param tests. DO NOT EDIT.

package unit_test

import (
	"encoding/json"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateIAM2TicketFlowParam_MarshalJSON(t *testing.T) {
	p := param.UpdateIAM2TicketFlowParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestUpdateIAM2TicketFlowParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.UpdateIAM2TicketFlowParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

func TestAddIAM2TicketFlowParam_MarshalJSON(t *testing.T) {
	p := param.AddIAM2TicketFlowParam{}
	data, err := json.Marshal(p)
	assertNoError(t, err)
	if len(data) == 0 {
		t.Fatal("marshaled JSON should not be empty")
	}
	// Verify it's valid JSON
	var raw map[string]interface{}
	assertNoError(t, json.Unmarshal(data, &raw))
}

func TestAddIAM2TicketFlowParam_UnmarshalJSON(t *testing.T) {
	jsonStr := `{}`
	var p param.AddIAM2TicketFlowParam
	err := json.Unmarshal([]byte(jsonStr), &p)
	assertNoError(t, err)
}

