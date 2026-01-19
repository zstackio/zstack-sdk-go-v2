// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNetworkServiceProvider(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryNetworkServiceProvider(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNetworkServiceProvider error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryNetworkServiceProvider result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageNetworkServiceProvider(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageNetworkServiceProvider(&queryParam)
	if err != nil {
		t.Errorf("TestPageNetworkServiceProvider error: %v", err)
		return
	}
	golog.Infof("PageNetworkServiceProvider result: total=%d, returned=%d", total, len(result))
}

func TestGetNetworkServiceProvider(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryNetworkServiceProvider(&queryParam)
	if err != nil {
		t.Errorf("TestGetNetworkServiceProvider Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NetworkServiceProvider found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetNetworkServiceProvider(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNetworkServiceProvider error: %v", err)
		return
	}
	golog.Infof("GetNetworkServiceProvider result: %s, Name: %s", result.UUID, result.Name)
}
