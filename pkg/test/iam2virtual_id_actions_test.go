// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualID(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2VirtualID(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualID error: %v", err)
		return
	}
	golog.Infof("QueryIAM2VirtualID result count: %d", len(result))
}
func TestGetIAM2VirtualID(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualID(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualID Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualID found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2VirtualID(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualID error: %v", err)
		return
	}
	golog.Infof("GetIAM2VirtualID result: %s", result.UUID)
}

func TestUpdateIAM2VirtualID(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualID(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualID Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualID found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2VirtualIDParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2VirtualIDParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2VirtualID(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualID error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2VirtualID result: %s", result.UUID)
}

func TestDeleteIAM2VirtualID(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIAM2VirtualID is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualID(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIAM2VirtualID Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualID found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIAM2VirtualID(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIAM2VirtualID error: %v", err)
		return
	}
	golog.Infof("DeleteIAM2VirtualID succeeded for UUID: %s", list[0].UUID)
}

func TestCreateIAM2VirtualID(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateIAM2VirtualID is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateIAM2VirtualIDParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateIAM2VirtualIDParamDetail{
	// 		Name: "test-iam2virtualid",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateIAM2VirtualID(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateIAM2VirtualID error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateIAM2VirtualID result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteIAM2VirtualID(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteIAM2VirtualID error: %v", err)
	// }
}

func TestLoginIAM2VirtualID(t *testing.T) {
	// LoginIAM2VirtualID operation
	t.Skip("TestLoginIAM2VirtualID requires manual implementation")

}
