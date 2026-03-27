// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateCephPrimaryStorageMon(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephPrimaryStorage(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestUpdateCephPrimaryStorageMon Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephPrimaryStorageMon found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateCephPrimaryStorageMonParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateCephPrimaryStorageMonParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCephPrimaryStorageMon(context.Background(), list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCephPrimaryStorageMon error: %v", err)
		return
	}
	golog.Infof("UpdateCephPrimaryStorageMon result: %s", result.UUID)
}
