// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2Organization(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2Organization error: %v", err)
		return
	}
	golog.Infof("QueryIAM2Organization result count: %d", len(result))
}

func TestUpdateIAM2Organization(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2Organization Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2Organization found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2OrganizationParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2OrganizationParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2Organization(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2Organization error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2Organization result: %s", result.UUID)
}

func TestDeleteIAM2Organization(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIAM2Organization is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIAM2Organization Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2Organization found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIAM2Organization(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIAM2Organization error: %v", err)
		return
	}
	golog.Infof("DeleteIAM2Organization succeeded for UUID: %s", list[0].UUID)
}

func TestCreateIAM2Organization(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateIAM2Organization is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateIAM2OrganizationParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateIAM2OrganizationParamDetail{
	// 		Name: "test-iam2organization",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateIAM2Organization(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateIAM2Organization error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateIAM2Organization result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteIAM2Organization(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteIAM2Organization error: %v", err)
	// }
}
