// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVni(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVniRange(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVni error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVniRange result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d-%d", r.UUID, r.Name, r.StartVni, r.EndVni)
	}
	golog.Infof("======================================")
}

func TestPageVniRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVniRange(&queryParam)
	if err != nil {
		t.Errorf("TestPageVniRange error: %v", err)
		return
	}
	golog.Infof("PageVniRange result: total=%d, returned=%d", total, len(result))
}

func TestGetVniRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVniRange(&queryParam)
	if err != nil {
		t.Errorf("TestGetVniRange Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VniRange found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVniRange(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVniRange error: %v", err)
		return
	}
	golog.Infof("GetVniRange result: %s, Name: %s", result.UUID, result.Name)
}
