// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryActiveAlarmTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryActiveAlarmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryActiveAlarmTemplate error: %v", err)
		return
	}
	golog.Infof("QueryActiveAlarmTemplate result count: %d", len(result))
}

func TestUpdateActiveAlarmTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryActiveAlarmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateActiveAlarmTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ActiveAlarmTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateActiveAlarmTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateActiveAlarmTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateActiveAlarmTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateActiveAlarmTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateActiveAlarmTemplate result: %s", result.UUID)
}
