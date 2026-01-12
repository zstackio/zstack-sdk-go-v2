// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateKVMHost(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryKvmHypervisorInfo(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateKVMHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No KVMHost found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateKVMHostParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateKVMHostParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateKVMHost(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateKVMHost error: %v", err)
		return
	}
	golog.Infof("UpdateKVMHost result: %s", result.UUID)
}

func TestAddKVMHost(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddKVMHost requires valid creation parameters")

}
