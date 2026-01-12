// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNasFileSystem(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNasFileSystem(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNasFileSystem error: %v", err)
		return
	}
	golog.Infof("QueryNasFileSystem result count: %d", len(result))
}
func TestGetNasFileSystem(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNasFileSystem(&queryParam)
	if err != nil {
		t.Errorf("TestGetNasFileSystem Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NasFileSystem found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNasFileSystem(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNasFileSystem error: %v", err)
		return
	}
	golog.Infof("GetNasFileSystem result: %s", result.UUID)
}

func TestUpdateNasFileSystem(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNasFileSystem(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateNasFileSystem Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NasFileSystem found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateNasFileSystemParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateNasFileSystemParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateNasFileSystem(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateNasFileSystem error: %v", err)
		return
	}
	golog.Infof("UpdateNasFileSystem result: %s", result.UUID)
}

func TestDeleteNasFileSystem(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteNasFileSystem is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNasFileSystem(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteNasFileSystem Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NasFileSystem found to test Delete")
		return
	}

	err = accountLoginCli.DeleteNasFileSystem(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteNasFileSystem error: %v", err)
		return
	}
	golog.Infof("DeleteNasFileSystem succeeded for UUID: %s", list[0].UUID)
}
