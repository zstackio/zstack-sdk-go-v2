// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("QueryMonitorTemplate result count: %d", len(result))
}

func TestUpdateMonitorTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateMonitorTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateMonitorTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateMonitorTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateMonitorTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateMonitorTemplate result: %s", result.UUID)
}

func TestDeleteMonitorTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMonitorTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMonitorTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMonitorTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteMonitorTemplate succeeded for UUID: %s", list[0].UUID)
}

func TestCreateMonitorTemplate(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateMonitorTemplate is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateMonitorTemplateParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateMonitorTemplateParamDetail{
	// 		Name: "test-monitortemplate",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateMonitorTemplate(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateMonitorTemplate error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateMonitorTemplate result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteMonitorTemplate(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteMonitorTemplate error: %v", err)
	// }
}

func TestCloneMonitorTemplate(t *testing.T) {
	// Clone operation
	t.Skip("TestCloneMonitorTemplate requires a valid resource to clone")

}
