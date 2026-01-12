// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSTextTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSTextTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSTextTemplate error: %v", err)
		return
	}
	golog.Infof("QuerySNSTextTemplate result count: %d", len(result))
}

func TestUpdateSNSTextTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSTextTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSTextTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSTextTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSTextTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSTextTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSTextTemplate(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSTextTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateSNSTextTemplate result: %s", result.Uuid)
}

func TestDeleteSNSTextTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSNSTextTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSTextTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSNSTextTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSTextTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSNSTextTemplate(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSNSTextTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteSNSTextTemplate succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateSNSTextTemplate(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSTextTemplate is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSTextTemplateParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSTextTemplateParamDetail{
	// 		Name: "test-snstexttemplate",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSTextTemplate(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSTextTemplate error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSTextTemplate result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSTextTemplate(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSTextTemplate error: %v", err)
	// }
}
