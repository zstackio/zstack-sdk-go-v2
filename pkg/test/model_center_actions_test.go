// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelCenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelCenter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelCenter error: %v", err)
		return
	}
	golog.Infof("QueryModelCenter result count: %d", len(result))
}

func TestUpdateModelCenter(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelCenter(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateModelCenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelCenter found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateModelCenterParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateModelCenterParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateModelCenter(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateModelCenter error: %v", err)
		return
	}
	golog.Infof("UpdateModelCenter result: %s", result.UUID)
}

func TestDeleteModelCenter(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteModelCenter is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelCenter(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteModelCenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelCenter found to test Delete")
		return
	}

	err = accountLoginCli.DeleteModelCenter(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteModelCenter error: %v", err)
		return
	}
	golog.Infof("DeleteModelCenter succeeded for UUID: %s", list[0].UUID)
}

func TestAddModelCenter(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddModelCenter requires valid creation parameters")

}
