// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageStoreBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageStoreBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageStoreBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryImageStoreBackupStorage result count: %d", len(result))
}

func TestUpdateImageStoreBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImageStoreBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateImageStoreBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImageStoreBackupStorage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateImageStoreBackupStorageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateImageStoreBackupStorageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateImageStoreBackupStorage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateImageStoreBackupStorage error: %v", err)
		return
	}
	golog.Infof("UpdateImageStoreBackupStorage result: %s", result.UUID)
}

func TestAddImageStoreBackupStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddImageStoreBackupStorage requires valid creation parameters")

}

func TestReconnectImageStoreBackupStorage(t *testing.T) {
	// ReconnectImageStoreBackupStorage operation
	t.Skip("TestReconnectImageStoreBackupStorage requires manual implementation")

}
