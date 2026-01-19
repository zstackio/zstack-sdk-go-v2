// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPriceTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPriceTable error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryPriceTable result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPagePriceTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PagePriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestPagePriceTable error: %v", err)
		return
	}
	golog.Infof("PagePriceTable result: total=%d, returned=%d", total, len(result))
}

func TestGetPriceTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestGetPriceTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PriceTable found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetPriceTable(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPriceTable error: %v", err)
		return
	}
	golog.Infof("GetPriceTable result: %s, Name: %s", result.UUID, result.Name)
}
