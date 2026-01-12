// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestVmScriptExecutedRecord(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestVmScriptExecutedRecord(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestVmScriptExecutedRecord error: %v", err)
		return
	}
	golog.Infof("QueryGuestVmScriptExecutedRecord result count: %d", len(result))
}
func TestGetGuestVmScriptExecutedRecord(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGuestVmScriptExecutedRecord(&queryParam)
	if err != nil {
		t.Errorf("TestGetGuestVmScriptExecutedRecord Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GuestVmScriptExecutedRecord found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetGuestVmScriptExecutedRecord(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetGuestVmScriptExecutedRecord error: %v", err)
		return
	}
	golog.Infof("GetGuestVmScriptExecutedRecord result: %s", result.UUID)
}
