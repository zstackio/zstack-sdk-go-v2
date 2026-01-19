// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcFirewall(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVpcFirewall(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcFirewall error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVpcFirewall result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageVpcFirewall(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVpcFirewall(&queryParam)
	if err != nil {
		t.Errorf("TestPageVpcFirewall error: %v", err)
		return
	}
	golog.Infof("PageVpcFirewall result: total=%d, returned=%d", total, len(result))
}

func TestGetVpcFirewall(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVpcFirewall(&queryParam)
	if err != nil {
		t.Errorf("TestGetVpcFirewall Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcFirewall found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVpcFirewall(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVpcFirewall error: %v", err)
		return
	}
	golog.Infof("GetVpcFirewall result: %s, Name: %s", result.UUID, result.Name)
}
