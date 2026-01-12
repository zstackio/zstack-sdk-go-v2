// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingRule error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingRule result count: %d", len(result))
}
func TestGetAutoScalingRule(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingRule(&queryParam)
	if err != nil {
		t.Errorf("TestGetAutoScalingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingRule found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAutoScalingRule(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAutoScalingRule error: %v", err)
		return
	}
	golog.Infof("GetAutoScalingRule result: %s", result.UUID)
}

func TestUpdateAutoScalingRule(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingRule(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingRule found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAutoScalingRuleParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAutoScalingRuleParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAutoScalingRule(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingRule error: %v", err)
		return
	}
	golog.Infof("UpdateAutoScalingRule result: %s", result.UUID)
}

func TestDeleteAutoScalingRule(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAutoScalingRule is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingRule(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingRule found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAutoScalingRule(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingRule error: %v", err)
		return
	}
	golog.Infof("DeleteAutoScalingRule succeeded for UUID: %s", list[0].UUID)
}
