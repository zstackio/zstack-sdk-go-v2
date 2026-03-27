// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateEcsVpc(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEcsVpcFromLocal(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestUpdateEcsVpc Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EcsVpc found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEcsVpcParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEcsVpcParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEcsVpc(context.Background(), list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEcsVpc error: %v", err)
		return
	}
	golog.Infof("UpdateEcsVpc result: %s", result.UUID)
}
