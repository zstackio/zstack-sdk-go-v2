// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingVmTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingVmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingVmTemplate error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingVmTemplate result count: %d", len(result))
}
func TestGetAutoScalingVmTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingVmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestGetAutoScalingVmTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingVmTemplate found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAutoScalingVmTemplate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAutoScalingVmTemplate error: %v", err)
		return
	}
	golog.Infof("GetAutoScalingVmTemplate result: %s", result.UUID)
}

func TestUpdateAutoScalingVmTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingVmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingVmTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingVmTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAutoScalingVmTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAutoScalingVmTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAutoScalingVmTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAutoScalingVmTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateAutoScalingVmTemplate result: %s", result.UUID)
}

func TestCreateAutoScalingVmTemplate(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAutoScalingVmTemplate is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAutoScalingVmTemplateParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAutoScalingVmTemplateParamDetail{
	// 		Name: "test-autoscalingvmtemplate",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAutoScalingVmTemplate(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAutoScalingVmTemplate error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAutoScalingVmTemplate result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAutoScalingVmTemplate(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAutoScalingVmTemplate error: %v", err)
	// }
}
