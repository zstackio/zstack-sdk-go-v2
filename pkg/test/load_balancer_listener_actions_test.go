// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancerListener(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryLoadBalancerListener result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%d", r.UUID, r.Name, r.Protocol, r.LoadBalancerPort)
	}
	golog.Infof("======================================")
}

func TestPageLoadBalancerListener(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestPageLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("PageLoadBalancerListener result: total=%d, returned=%d", total, len(result))
}

func TestGetLoadBalancerListener(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestGetLoadBalancerListener Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerListener found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetLoadBalancerListener(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("GetLoadBalancerListener result: %s, Name: %s", result.UUID, result.Name)
}
