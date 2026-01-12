// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelService(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelService(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelService error: %v", err)
		return
	}
	golog.Infof("QueryModelService result count: %d", len(result))
}
func TestGetModelService(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelService(&queryParam)
	if err != nil {
		t.Errorf("TestGetModelService Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelService found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetModelService(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetModelService error: %v", err)
		return
	}
	golog.Infof("GetModelService result: %s", result.UUID)
}

func TestUpdateModelService(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelService(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateModelService Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelService found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateModelServiceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateModelServiceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateModelService(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateModelService error: %v", err)
		return
	}
	golog.Infof("UpdateModelService result: %s", result.UUID)
}

func TestDeleteModelService(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteModelService is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelService(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteModelService Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelService found to test Delete")
		return
	}

	err = accountLoginCli.DeleteModelService(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteModelService error: %v", err)
		return
	}
	golog.Infof("DeleteModelService succeeded for UUID: %s", list[0].UUID)
}

func TestCloneModelService(t *testing.T) {
	// Clone operation
	t.Skip("TestCloneModelService requires a valid resource to clone")

}

func TestAddModelService(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddModelService requires valid creation parameters")

}
