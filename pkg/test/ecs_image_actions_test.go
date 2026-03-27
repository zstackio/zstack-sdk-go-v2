// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateEcsImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEcsImageFromLocal(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestUpdateEcsImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EcsImage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEcsImageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEcsImageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEcsImage(context.Background(), list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEcsImage error: %v", err)
		return
	}
	golog.Infof("UpdateEcsImage result: %s", result.UUID)
}
