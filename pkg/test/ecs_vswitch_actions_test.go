// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateEcsVSwitch(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEcsVSwitchFromLocal(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestUpdateEcsVSwitch Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EcsVSwitch found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEcsVSwitchParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEcsVSwitchParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEcsVSwitch(context.Background(), list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEcsVSwitch error: %v", err)
		return
	}
	golog.Infof("UpdateEcsVSwitch result: %s", result.UUID)
}
