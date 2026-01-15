// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAlertDataAck(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAlertDataAck(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAlertDataAck error: %v", err)
		return
	}
	golog.Infof("QueryAlertDataAck result count: %d", len(result))
}
func TestGetAlertDataAck(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlertDataAck(&queryParam)
	if err != nil {
		t.Errorf("TestGetAlertDataAck Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AlertDataAck found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAlertDataAck(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAlertDataAck error: %v", err)
		return
	}
	golog.Infof("GetAlertDataAck result: %s", result.UUID)
}

func TestUpdateAlertDataAck(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlertDataAck(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAlertDataAck Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AlertDataAck found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAlertDataAckParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAlertDataAckParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAlertDataAck(list[0].AlertDataUuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAlertDataAck error: %v", err)
		return
	}
	golog.Infof("UpdateAlertDataAck result: %s", result.AlertDataUuid)
}
