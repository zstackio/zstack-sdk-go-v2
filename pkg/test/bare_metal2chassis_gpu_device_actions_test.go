// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2ChassisGpuDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2ChassisGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2ChassisGpuDevice error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2ChassisGpuDevice result count: %d", len(result))
}
func TestGetBareMetal2ChassisGpuDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ChassisGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetBareMetal2ChassisGpuDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ChassisGpuDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBareMetal2ChassisGpuDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBareMetal2ChassisGpuDevice error: %v", err)
		return
	}
	golog.Infof("GetBareMetal2ChassisGpuDevice result: %s", result.UUID)
}
