// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostNetworkBonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryHostNetworkBonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostNetworkBonding error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryHostNetworkBonding result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.BondingName)
	}
	golog.Infof("======================================")
}

func TestPageHostNetworkBonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageHostNetworkBonding(&queryParam)
	if err != nil {
		t.Errorf("TestPageHostNetworkBonding error: %v", err)
		return
	}
	golog.Infof("PageHostNetworkBonding result: total=%d, returned=%d", total, len(result))
}

func TestGetHostNetworkBonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryHostNetworkBonding(&queryParam)
	if err != nil {
		t.Errorf("TestGetHostNetworkBonding Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostNetworkBonding found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetHostNetworkBonding(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHostNetworkBonding error: %v", err)
		return
	}
	golog.Infof("GetHostNetworkBonding result: %s, Name: %s", result.UUID, result.BondingName)
}
