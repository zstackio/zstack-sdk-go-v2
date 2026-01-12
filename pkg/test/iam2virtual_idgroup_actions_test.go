// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualIDGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("QueryIAM2VirtualIDGroup result count: %d", len(result))
}
func TestGetIAM2VirtualIDGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualIDGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2VirtualIDGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("GetIAM2VirtualIDGroup result: %s", result.UUID)
}

func TestUpdateIAM2VirtualIDGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualIDGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2VirtualIDGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2VirtualIDGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2VirtualIDGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2VirtualIDGroup result: %s", result.UUID)
}

func TestDeleteIAM2VirtualIDGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIAM2VirtualIDGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIAM2VirtualIDGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIAM2VirtualIDGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("DeleteIAM2VirtualIDGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateIAM2VirtualIDGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateIAM2VirtualIDGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateIAM2VirtualIDGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateIAM2VirtualIDGroupParamDetail{
	// 		Name: "test-iam2virtualidgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateIAM2VirtualIDGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateIAM2VirtualIDGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateIAM2VirtualIDGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteIAM2VirtualIDGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteIAM2VirtualIDGroup error: %v", err)
	// }
}
