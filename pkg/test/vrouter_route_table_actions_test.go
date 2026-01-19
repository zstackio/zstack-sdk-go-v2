// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVRouterRouteTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVRouterRouteTable result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageVRouterRouteTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestPageVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("PageVRouterRouteTable result: total=%d, returned=%d", total, len(result))
}

func TestGetVRouterRouteTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVRouterRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestGetVRouterRouteTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VRouterRouteTable found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVRouterRouteTable(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVRouterRouteTable error: %v", err)
		return
	}
	golog.Infof("GetVRouterRouteTable result: %s, Name: %s", result.UUID, result.Name)
}
