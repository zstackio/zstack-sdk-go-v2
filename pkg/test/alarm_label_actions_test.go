// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateAlarmLabel(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAlarmLabel Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AlarmLabel found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAlarmLabelParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAlarmLabelParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAlarmLabel(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAlarmLabel error: %v", err)
		return
	}
	golog.Infof("UpdateAlarmLabel result: %s", result.UUID)
}
