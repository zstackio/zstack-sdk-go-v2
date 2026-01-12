// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmSchedulingRuleGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("QueryVmSchedulingRuleGroup result count: %d", len(result))
}
func TestGetVmSchedulingRuleGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmSchedulingRuleGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmSchedulingRuleGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmSchedulingRuleGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("GetVmSchedulingRuleGroup result: %s", result.UUID)
}

func TestUpdateVmSchedulingRuleGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVmSchedulingRuleGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmSchedulingRuleGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVmSchedulingRuleGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVmSchedulingRuleGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVmSchedulingRuleGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVmSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("UpdateVmSchedulingRuleGroup result: %s", result.UUID)
}

func TestDeleteVmSchedulingRuleGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVmSchedulingRuleGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVmSchedulingRuleGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmSchedulingRuleGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVmSchedulingRuleGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVmSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("DeleteVmSchedulingRuleGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVmSchedulingRuleGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVmSchedulingRuleGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVmSchedulingRuleGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVmSchedulingRuleGroupParamDetail{
	// 		Name: "test-vmschedulingrulegroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVmSchedulingRuleGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVmSchedulingRuleGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVmSchedulingRuleGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVmSchedulingRuleGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVmSchedulingRuleGroup error: %v", err)
	// }
}
