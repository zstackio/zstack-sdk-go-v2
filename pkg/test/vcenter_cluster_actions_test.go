// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVCenterCluster(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterCluster error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVCenterCluster result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageVCenterCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVCenterCluster(&queryParam)
	if err != nil {
		t.Errorf("TestPageVCenterCluster error: %v", err)
		return
	}
	golog.Infof("PageVCenterCluster result: total=%d, returned=%d", total, len(result))
}

func TestGetVCenterCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVCenterCluster(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterCluster found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVCenterCluster(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterCluster error: %v", err)
		return
	}
	golog.Infof("GetVCenterCluster result: %s, Name: %s", result.UUID, result.Name)
}
