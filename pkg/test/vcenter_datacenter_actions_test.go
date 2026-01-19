// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterDatacenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVCenterDatacenter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterDatacenter error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVCenterDatacenter result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageVCenterDatacenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVCenterDatacenter(&queryParam)
	if err != nil {
		t.Errorf("TestPageVCenterDatacenter error: %v", err)
		return
	}
	golog.Infof("PageVCenterDatacenter result: total=%d, returned=%d", total, len(result))
}

func TestGetVCenterDatacenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVCenterDatacenter(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterDatacenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterDatacenter found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVCenterDatacenter(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterDatacenter error: %v", err)
		return
	}
	golog.Infof("GetVCenterDatacenter result: %s, Name: %s", result.UUID, result.Name)
}
