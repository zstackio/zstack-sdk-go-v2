// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteRuleSet(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteRuleSet(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteRuleSet error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteRuleSet result count: %d", len(result))
}

func TestUpdatePolicyRouteRuleSet(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteRuleSet(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePolicyRouteRuleSet Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteRuleSet found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePolicyRouteRuleSetParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePolicyRouteRuleSetParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePolicyRouteRuleSet(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePolicyRouteRuleSet error: %v", err)
		return
	}
	golog.Infof("UpdatePolicyRouteRuleSet result: %s", result.UUID)
}

func TestDeletePolicyRouteRuleSet(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePolicyRouteRuleSet is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteRuleSet(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteRuleSet Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteRuleSet found to test Delete")
		return
	}

	err = accountLoginCli.DeletePolicyRouteRuleSet(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteRuleSet error: %v", err)
		return
	}
	golog.Infof("DeletePolicyRouteRuleSet succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePolicyRouteRuleSet(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePolicyRouteRuleSet is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePolicyRouteRuleSetParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePolicyRouteRuleSetParamDetail{
	// 		Name: "test-policyrouteruleset",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePolicyRouteRuleSet(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePolicyRouteRuleSet error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePolicyRouteRuleSet result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePolicyRouteRuleSet(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePolicyRouteRuleSet error: %v", err)
	// }
}
