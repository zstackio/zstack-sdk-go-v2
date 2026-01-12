// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2Project(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2Project(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2Project error: %v", err)
		return
	}
	golog.Infof("QueryIAM2Project result count: %d", len(result))
}

func TestUpdateIAM2Project(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2Project(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2Project Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2Project found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2ProjectParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2ProjectParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2Project(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2Project error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2Project result: %s", result.UUID)
}

func TestDeleteIAM2Project(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIAM2Project is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2Project(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIAM2Project Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2Project found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIAM2Project(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIAM2Project error: %v", err)
		return
	}
	golog.Infof("DeleteIAM2Project succeeded for UUID: %s", list[0].UUID)
}

func TestCreateIAM2Project(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateIAM2Project is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateIAM2ProjectParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateIAM2ProjectParamDetail{
	// 		Name: "test-iam2project",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateIAM2Project(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateIAM2Project error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateIAM2Project result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteIAM2Project(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteIAM2Project error: %v", err)
	// }
}

func TestRecoverIAM2Project(t *testing.T) {
	// Recover operation - requires a deleted resource
	t.Skip("TestRecoverIAM2Project requires a deleted resource UUID")

}

func TestExpungeIAM2Project(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeIAM2Project is dangerous - permanently deletes resource")

}

func TestLoginIAM2Project(t *testing.T) {
	// LoginIAM2Project operation
	t.Skip("TestLoginIAM2Project requires manual implementation")

}
