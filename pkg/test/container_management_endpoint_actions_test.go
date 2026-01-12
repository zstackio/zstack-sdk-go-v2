// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryContainerManagementEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryContainerManagementEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQueryContainerManagementEndpoint error: %v", err)
		return
	}
	golog.Infof("QueryContainerManagementEndpoint result count: %d", len(result))
}
func TestGetContainerManagementEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryContainerManagementEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestGetContainerManagementEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ContainerManagementEndpoint found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetContainerManagementEndpoint(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetContainerManagementEndpoint error: %v", err)
		return
	}
	golog.Infof("GetContainerManagementEndpoint result: %s", result.UUID)
}

func TestUpdateContainerManagementEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryContainerManagementEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateContainerManagementEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ContainerManagementEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateContainerManagementEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateContainerManagementEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateContainerManagementEndpoint(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateContainerManagementEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateContainerManagementEndpoint result: %s", result.UUID)
}

func TestDeleteContainerManagementEndpoint(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteContainerManagementEndpoint is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryContainerManagementEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteContainerManagementEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ContainerManagementEndpoint found to test Delete")
		return
	}

	err = accountLoginCli.DeleteContainerManagementEndpoint(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteContainerManagementEndpoint error: %v", err)
		return
	}
	golog.Infof("DeleteContainerManagementEndpoint succeeded for UUID: %s", list[0].UUID)
}

func TestAddContainerManagementEndpoint(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddContainerManagementEndpoint requires valid creation parameters")

}

func TestSyncContainerManagementEndpoint(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncContainerManagementEndpoint requires a valid resource to sync")

}
