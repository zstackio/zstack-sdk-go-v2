// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestVmScriptExecutedRecordDetail(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestVmScriptExecutedRecordDetail(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestVmScriptExecutedRecordDetail error: %v", err)
		return
	}
	golog.Infof("QueryGuestVmScriptExecutedRecordDetail result count: %d", len(result))
}
func TestGetGuestVmScriptExecutedRecordDetail(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGuestVmScriptExecutedRecordDetail(&queryParam)
	if err != nil {
		t.Errorf("TestGetGuestVmScriptExecutedRecordDetail Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GuestVmScriptExecutedRecordDetail found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetGuestVmScriptExecutedRecordDetail(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetGuestVmScriptExecutedRecordDetail error: %v", err)
		return
	}
	golog.Infof("GetGuestVmScriptExecutedRecordDetail result: %s", result.UUID)
}
