// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBlockPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBlockPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBlockPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryBlockPrimaryStorage result count: %d", len(result))
}
func TestGetBlockPrimaryStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBlockPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetBlockPrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BlockPrimaryStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBlockPrimaryStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBlockPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("GetBlockPrimaryStorage result: %s", result.UUID)
}

func TestUpdateBlockPrimaryStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBlockPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBlockPrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BlockPrimaryStorage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBlockPrimaryStorageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBlockPrimaryStorageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBlockPrimaryStorage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBlockPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("UpdateBlockPrimaryStorage result: %s", result.UUID)
}

func TestAddBlockPrimaryStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddBlockPrimaryStorage requires valid creation parameters")

}
