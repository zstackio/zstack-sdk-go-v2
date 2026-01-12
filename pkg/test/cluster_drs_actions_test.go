// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryClusterDRS(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryClusterDRS(&queryParam)
	if err != nil {
		t.Errorf("TestQueryClusterDRS error: %v", err)
		return
	}
	golog.Infof("QueryClusterDRS result count: %d", len(result))
}

func TestUpdateClusterDRS(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryClusterDRS(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateClusterDRS Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ClusterDRS found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateClusterDRSParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateClusterDRSParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateClusterDRS(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateClusterDRS error: %v", err)
		return
	}
	golog.Infof("UpdateClusterDRS result: %s", result.UUID)
}

func TestDeleteClusterDRS(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteClusterDRS is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryClusterDRS(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteClusterDRS Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ClusterDRS found to test Delete")
		return
	}

	err = accountLoginCli.DeleteClusterDRS(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteClusterDRS error: %v", err)
		return
	}
	golog.Infof("DeleteClusterDRS succeeded for UUID: %s", list[0].UUID)
}

func TestCreateClusterDRS(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateClusterDRS is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateClusterDRSParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateClusterDRSParamDetail{
	// 		Name: "test-clusterdrs",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateClusterDRS(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateClusterDRS error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateClusterDRS result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteClusterDRS(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteClusterDRS error: %v", err)
	// }
}
