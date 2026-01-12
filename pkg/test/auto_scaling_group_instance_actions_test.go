// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingGroupInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingGroupInstance error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingGroupInstance result count: %d", len(result))
}
func TestGetAutoScalingGroupInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestGetAutoScalingGroupInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroupInstance found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAutoScalingGroupInstance(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAutoScalingGroupInstance error: %v", err)
		return
	}
	golog.Infof("GetAutoScalingGroupInstance result: %s", result.UUID)
}

func TestUpdateAutoScalingGroupInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingGroupInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroupInstance found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAutoScalingGroupInstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAutoScalingGroupInstanceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAutoScalingGroupInstance(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingGroupInstance error: %v", err)
		return
	}
	golog.Infof("UpdateAutoScalingGroupInstance result: %s", result.UUID)
}

func TestDeleteAutoScalingGroupInstance(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAutoScalingGroupInstance is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingGroupInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroupInstance found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAutoScalingGroupInstance(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingGroupInstance error: %v", err)
		return
	}
	golog.Infof("DeleteAutoScalingGroupInstance succeeded for UUID: %s", list[0].UUID)
}
