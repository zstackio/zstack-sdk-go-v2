// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2ProjectTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectTemplate error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectTemplate result count: %d", len(result))
}
func TestGetIAM2ProjectTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2ProjectTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2ProjectTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2ProjectTemplate found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2ProjectTemplate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2ProjectTemplate error: %v", err)
		return
	}
	golog.Infof("GetIAM2ProjectTemplate result: %s", result.UUID)
}

func TestUpdateIAM2ProjectTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2ProjectTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2ProjectTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2ProjectTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2ProjectTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2ProjectTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2ProjectTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2ProjectTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2ProjectTemplate result: %s", result.UUID)
}

func TestDeleteIAM2ProjectTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIAM2ProjectTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2ProjectTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIAM2ProjectTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2ProjectTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIAM2ProjectTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIAM2ProjectTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteIAM2ProjectTemplate succeeded for UUID: %s", list[0].UUID)
}

func TestCreateIAM2ProjectTemplate(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateIAM2ProjectTemplate is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateIAM2ProjectTemplateParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateIAM2ProjectTemplateParamDetail{
	// 		Name: "test-iam2projecttemplate",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateIAM2ProjectTemplate(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateIAM2ProjectTemplate error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateIAM2ProjectTemplate result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteIAM2ProjectTemplate(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteIAM2ProjectTemplate error: %v", err)
	// }
}
