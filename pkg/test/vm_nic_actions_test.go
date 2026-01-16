// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmNic(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVmNic(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmNic error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVmNic result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Mac, r.L3NetworkUuid, r.VmInstanceUuid)
	}
	golog.Infof("======================================")
}

func TestPageVmNic(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVmNic(&queryParam)
	if err != nil {
		t.Errorf("TestPageVmNic error: %v", err)
		return
	}
	golog.Infof("PageVmNic result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Mac, r.L3NetworkUuid)
	}
}

func TestGetVmNic(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVmNic(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmNic Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmNic found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVmNic(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmNic error: %v", err)
		return
	}
	golog.Infof("GetVmNic result: %s, MAC: %s", result.UUID, result.Mac)
}
