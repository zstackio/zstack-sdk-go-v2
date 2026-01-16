// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIpRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIpRange error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryIpRange result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.StartIp, r.EndIp)
	}
	golog.Infof("======================================")
}

func TestPageIpRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestPageIpRange error: %v", err)
		return
	}
	golog.Infof("PageIpRange result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s-%s", r.UUID, r.Name, r.StartIp, r.EndIp)
	}
}

func TestGetIpRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestGetIpRange Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IpRange found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetIpRange(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIpRange error: %v", err)
		return
	}
	golog.Infof("GetIpRange result: %s, Name: %s, Range: %s-%s", result.UUID, result.Name, result.StartIp, result.EndIp)
}
