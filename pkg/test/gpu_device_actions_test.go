// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGpuDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGpuDevice error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryGpuDevice result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageGpuDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestPageGpuDevice error: %v", err)
		return
	}
	golog.Infof("PageGpuDevice result: total=%d, returned=%d", total, len(result))
}

func TestGetGpuDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryGpuDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetGpuDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GpuDevice found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetGpuDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetGpuDevice error: %v", err)
		return
	}
	golog.Infof("GetGpuDevice result: %s, Name: %s", result.UUID, result.Name)
}
