// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPciDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDevice error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryPciDevice result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPagePciDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PagePciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestPagePciDevice error: %v", err)
		return
	}
	golog.Infof("PagePciDevice result: total=%d, returned=%d", total, len(result))
}

func TestGetPciDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetPciDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PciDevice found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetPciDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPciDevice error: %v", err)
		return
	}
	golog.Infof("GetPciDevice result: %s, Name: %s", result.UUID, result.Name)
}
