// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPortForwardingRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryPortForwardingRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPortForwardingRule error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryPortForwardingRule result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%d", r.UUID, r.Name, r.VipIp, r.VipPortStart)
	}
	golog.Infof("======================================")
}

func TestPagePortForwardingRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PagePortForwardingRule(&queryParam)
	if err != nil {
		t.Errorf("TestPagePortForwardingRule error: %v", err)
		return
	}
	golog.Infof("PagePortForwardingRule result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.VipIp)
	}
}

func TestGetPortForwardingRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryPortForwardingRule(&queryParam)
	if err != nil {
		t.Errorf("TestGetPortForwardingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortForwardingRule found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetPortForwardingRule(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPortForwardingRule error: %v", err)
		return
	}
	golog.Infof("GetPortForwardingRule result: %s, Name: %s", result.UUID, result.Name)
}
