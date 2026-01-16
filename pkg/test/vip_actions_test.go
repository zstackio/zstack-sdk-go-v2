// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVip(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVip(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVip error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVip result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.Ip, r.State)
	}
	golog.Infof("======================================")
}

func TestPageVip(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVip(&queryParam)
	if err != nil {
		t.Errorf("TestPageVip error: %v", err)
		return
	}
	golog.Infof("PageVip result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Ip)
	}
}

func TestGetVip(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVip(&queryParam)
	if err != nil {
		t.Errorf("TestGetVip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Vip found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVip(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVip error: %v", err)
		return
	}
	golog.Infof("GetVip result: %s, Name: %s, IP: %s", result.UUID, result.Name, result.Ip)
}
