// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVirtualRouterOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVirtualRouterOffering result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageVirtualRouterOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVirtualRouterOffering(&queryParam)
	if err != nil {
		t.Errorf("TestPageVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("PageVirtualRouterOffering result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
}

func TestGetVirtualRouterOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVirtualRouterOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetVirtualRouterOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VirtualRouterOffering found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVirtualRouterOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("GetVirtualRouterOffering result: %s, Name: %s", result.UUID, result.Name)
}
