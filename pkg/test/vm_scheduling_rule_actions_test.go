// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmSchedulingRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmSchedulingRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmSchedulingRule error: %v", err)
		return
	}
	golog.Infof("QueryVmSchedulingRule result count: %d", len(result))
}

func TestUpdateVmSchedulingRule(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmSchedulingRule(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVmSchedulingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmSchedulingRule found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVmSchedulingRuleParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVmSchedulingRuleParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVmSchedulingRule(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVmSchedulingRule error: %v", err)
		return
	}
	golog.Infof("UpdateVmSchedulingRule result: %s", result.Uuid)
}

func TestCreateVmSchedulingRule(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVmSchedulingRule is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVmSchedulingRuleParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVmSchedulingRuleParamDetail{
	// 		Name: "test-vmschedulingrule",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVmSchedulingRule(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVmSchedulingRule error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVmSchedulingRule result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVmSchedulingRule(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVmSchedulingRule error: %v", err)
	// }
}

func TestRemoveVmSchedulingRule(t *testing.T) {
	// RemoveVmSchedulingRule operation
	t.Skip("TestRemoveVmSchedulingRule requires manual implementation")

}

func TestValidateVmSchedulingRule(t *testing.T) {
	// ValidateVmSchedulingRule operation
	t.Skip("TestValidateVmSchedulingRule requires manual implementation")

}
