// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGpuDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGpuDevice error: %v", err)
		return
	}
	golog.Infof("QueryGpuDevice result count: %d", len(result))
}
func TestGetGpuDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetGpuDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GpuDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetGpuDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetGpuDevice error: %v", err)
		return
	}
	golog.Infof("GetGpuDevice result: %s", result.UUID)
}
