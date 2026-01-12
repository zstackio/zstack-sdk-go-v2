// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateCephBackupStorageMon(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateCephBackupStorageMon Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephBackupStorageMon found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateCephBackupStorageMonParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateCephBackupStorageMonParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCephBackupStorageMon(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCephBackupStorageMon error: %v", err)
		return
	}
	golog.Infof("UpdateCephBackupStorageMon result: %s", result.UUID)
}
