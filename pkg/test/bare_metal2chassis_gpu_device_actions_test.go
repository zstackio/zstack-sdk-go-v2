// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

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

