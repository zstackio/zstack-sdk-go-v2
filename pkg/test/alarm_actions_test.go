// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAlarm error: %v", err)
		return
	}
	golog.Infof("QueryAlarm result count: %d", len(result))
}

func TestUpdateAlarm(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAlarm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Alarm found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAlarmParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAlarmParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAlarm(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAlarm error: %v", err)
		return
	}
	golog.Infof("UpdateAlarm result: %s", result.UUID)
}

func TestDeleteAlarm(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAlarm is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAlarm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Alarm found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAlarm(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAlarm error: %v", err)
		return
	}
	golog.Infof("DeleteAlarm succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAlarm(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAlarm is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAlarmParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAlarmParamDetail{
	// 		Name: "test-alarm",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAlarm(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAlarm error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAlarm result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAlarm(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAlarm error: %v", err)
	// }
}
