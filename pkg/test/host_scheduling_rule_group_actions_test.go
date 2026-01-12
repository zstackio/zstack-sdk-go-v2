// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostSchedulingRuleGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("QueryHostSchedulingRuleGroup result count: %d", len(result))
}

func TestUpdateHostSchedulingRuleGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateHostSchedulingRuleGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostSchedulingRuleGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateHostSchedulingRuleGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateHostSchedulingRuleGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateHostSchedulingRuleGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateHostSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("UpdateHostSchedulingRuleGroup result: %s", result.UUID)
}

func TestDeleteHostSchedulingRuleGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteHostSchedulingRuleGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostSchedulingRuleGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteHostSchedulingRuleGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostSchedulingRuleGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteHostSchedulingRuleGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteHostSchedulingRuleGroup error: %v", err)
		return
	}
	golog.Infof("DeleteHostSchedulingRuleGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateHostSchedulingRuleGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateHostSchedulingRuleGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateHostSchedulingRuleGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateHostSchedulingRuleGroupParamDetail{
	// 		Name: "test-hostschedulingrulegroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateHostSchedulingRuleGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateHostSchedulingRuleGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateHostSchedulingRuleGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteHostSchedulingRuleGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteHostSchedulingRuleGroup error: %v", err)
	// }
}
