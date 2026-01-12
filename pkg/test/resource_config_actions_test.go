// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryResourceConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryResourceConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryResourceConfig error: %v", err)
		return
	}
	golog.Infof("QueryResourceConfig result count: %d", len(result))
}

func TestUpdateResourceConfig(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryResourceConfig(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateResourceConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ResourceConfig found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateResourceConfigParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateResourceConfigParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateResourceConfig(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateResourceConfig error: %v", err)
		return
	}
	golog.Infof("UpdateResourceConfig result: %s", result.UUID)
}

func TestDeleteResourceConfig(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteResourceConfig is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryResourceConfig(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteResourceConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ResourceConfig found to test Delete")
		return
	}

	err = accountLoginCli.DeleteResourceConfig(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteResourceConfig error: %v", err)
		return
	}
	golog.Infof("DeleteResourceConfig succeeded for UUID: %s", list[0].UUID)
}

func TestGetResourceConfig(t *testing.T) {
	// GetResourceConfig operation
	t.Skip("TestGetResourceConfig requires manual implementation")

}
