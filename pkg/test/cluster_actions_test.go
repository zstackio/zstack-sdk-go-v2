// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCluster(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCluster error: %v", err)
		return
	}
	golog.Infof("QueryCluster result count: %d", len(result))
}
func TestGetCluster(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCluster(&queryParam)
	if err != nil {
		t.Errorf("TestGetCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Cluster found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCluster(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCluster error: %v", err)
		return
	}
	golog.Infof("GetCluster result: %s", result.UUID)
}

func TestUpdateCluster(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCluster(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Cluster found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateClusterParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateClusterParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCluster(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCluster error: %v", err)
		return
	}
	golog.Infof("UpdateCluster result: %s", result.UUID)
}

func TestDeleteCluster(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCluster is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCluster(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Cluster found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCluster(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCluster error: %v", err)
		return
	}
	golog.Infof("DeleteCluster succeeded for UUID: %s", list[0].UUID)
}

func TestCreateCluster(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateCluster is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateClusterParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateClusterParamDetail{
	// 		Name: "test-cluster",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateCluster(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateCluster error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateCluster result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteCluster(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteCluster error: %v", err)
	// }
}
