// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostNetworkInterface(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostNetworkInterface error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryHostNetworkInterface result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.InterfaceName, r.IpAddresses)
	}
	golog.Infof("======================================")
}

func TestPageHostNetworkInterface(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestPageHostNetworkInterface error: %v", err)
		return
	}
	golog.Infof("PageHostNetworkInterface result: total=%d, returned=%d", total, len(result))
}

func TestGetHostNetworkInterface(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryHostNetworkInterface(&queryParam)
	if err != nil {
		t.Errorf("TestGetHostNetworkInterface Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostNetworkInterface found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetHostNetworkInterface(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHostNetworkInterface error: %v", err)
		return
	}
	golog.Infof("GetHostNetworkInterface result: %s, Name: %s", result.UUID, result.InterfaceName)
}
