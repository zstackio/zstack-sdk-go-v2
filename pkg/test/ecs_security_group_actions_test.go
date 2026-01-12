// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateEcsSecurityGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEcsSecurityGroupFromLocal(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateEcsSecurityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EcsSecurityGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEcsSecurityGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEcsSecurityGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEcsSecurityGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEcsSecurityGroup error: %v", err)
		return
	}
	golog.Infof("UpdateEcsSecurityGroup result: %s", result.UUID)
}
