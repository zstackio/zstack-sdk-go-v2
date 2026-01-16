// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVirtualRouterVm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVirtualRouterVm result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageVirtualRouterVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVirtualRouterVm(&queryParam)
	if err != nil {
		t.Errorf("TestPageVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("PageVirtualRouterVm result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
}

func TestGetVirtualRouterVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVirtualRouterVm(&queryParam)
	if err != nil {
		t.Errorf("TestGetVirtualRouterVm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VirtualRouterVm found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVirtualRouterVm(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("GetVirtualRouterVm result: %s, Name: %s", result.UUID, result.Name)
}
