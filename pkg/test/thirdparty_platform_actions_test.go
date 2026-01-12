// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryThirdpartyPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryThirdpartyPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQueryThirdpartyPlatform error: %v", err)
		return
	}
	golog.Infof("QueryThirdpartyPlatform result count: %d", len(result))
}
func TestGetThirdpartyPlatform(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryThirdpartyPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestGetThirdpartyPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ThirdpartyPlatform found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetThirdpartyPlatform(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetThirdpartyPlatform error: %v", err)
		return
	}
	golog.Infof("GetThirdpartyPlatform result: %s", result.UUID)
}

func TestUpdateThirdpartyPlatform(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryThirdpartyPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateThirdpartyPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ThirdpartyPlatform found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateThirdpartyPlatformParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateThirdpartyPlatformParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateThirdpartyPlatform(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateThirdpartyPlatform error: %v", err)
		return
	}
	golog.Infof("UpdateThirdpartyPlatform result: %s", result.UUID)
}

func TestDeleteThirdpartyPlatform(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteThirdpartyPlatform is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryThirdpartyPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteThirdpartyPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ThirdpartyPlatform found to test Delete")
		return
	}

	err = accountLoginCli.DeleteThirdpartyPlatform(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteThirdpartyPlatform error: %v", err)
		return
	}
	golog.Infof("DeleteThirdpartyPlatform succeeded for UUID: %s", list[0].UUID)
}

func TestAddThirdpartyPlatform(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddThirdpartyPlatform requires valid creation parameters")

}
