// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingGroup error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingGroup result count: %d", len(result))
}

func TestUpdateAutoScalingGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAutoScalingGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAutoScalingGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAutoScalingGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingGroup error: %v", err)
		return
	}
	golog.Infof("UpdateAutoScalingGroup result: %s", result.UUID)
}

func TestDeleteAutoScalingGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAutoScalingGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAutoScalingGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingGroup error: %v", err)
		return
	}
	golog.Infof("DeleteAutoScalingGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAutoScalingGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAutoScalingGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAutoScalingGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAutoScalingGroupParamDetail{
	// 		Name: "test-autoscalinggroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAutoScalingGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAutoScalingGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAutoScalingGroup result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAutoScalingGroup(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAutoScalingGroup error: %v", err)
	// }
}
