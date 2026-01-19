// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VxlanNetworkPool(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryL2VxlanNetworkPool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VxlanNetworkPool error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryL2VxlanNetworkPool result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageL2VxlanNetworkPool(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageL2VxlanNetworkPool(&queryParam)
	if err != nil {
		t.Errorf("TestPageL2VxlanNetworkPool error: %v", err)
		return
	}
	golog.Infof("PageL2VxlanNetworkPool result: total=%d, returned=%d", total, len(result))
}

func TestGetL2VxlanNetworkPool(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryL2VxlanNetworkPool(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetworkPool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2VxlanNetworkPool found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetL2VxlanNetworkPool(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetworkPool error: %v", err)
		return
	}
	golog.Infof("GetL2VxlanNetworkPool result: %s, Name: %s", result.UUID, result.Name)
}
