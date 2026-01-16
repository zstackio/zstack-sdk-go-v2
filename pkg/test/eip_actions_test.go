// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEip(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEip error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryEip result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.VipIp, r.State)
	}
	golog.Infof("======================================")
}

func TestPageEip(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageEip(&queryParam)
	if err != nil {
		t.Errorf("TestPageEip error: %v", err)
		return
	}
	golog.Infof("PageEip result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.VipIp)
	}
}

func TestGetEip(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestGetEip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Eip found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetEip(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetEip error: %v", err)
		return
	}
	golog.Infof("GetEip result: %s, Name: %s, IP: %s", result.UUID, result.Name, result.VipIp)
}
