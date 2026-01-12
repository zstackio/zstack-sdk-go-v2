// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAppBuildSystem(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAppBuildSystem(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAppBuildSystem error: %v", err)
		return
	}
	golog.Infof("QueryAppBuildSystem result count: %d", len(result))
}

func TestUpdateAppBuildSystem(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAppBuildSystem(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAppBuildSystem Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AppBuildSystem found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAppBuildSystemParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAppBuildSystemParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAppBuildSystem(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAppBuildSystem error: %v", err)
		return
	}
	golog.Infof("UpdateAppBuildSystem result: %s", result.UUID)
}

func TestDeleteAppBuildSystem(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAppBuildSystem is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAppBuildSystem(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAppBuildSystem Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AppBuildSystem found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAppBuildSystem(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAppBuildSystem error: %v", err)
		return
	}
	golog.Infof("DeleteAppBuildSystem succeeded for UUID: %s", list[0].UUID)
}

func TestAddAppBuildSystem(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddAppBuildSystem requires valid creation parameters")

}

func TestReconnectAppBuildSystem(t *testing.T) {
	// ReconnectAppBuildSystem operation
	t.Skip("TestReconnectAppBuildSystem requires manual implementation")

}
