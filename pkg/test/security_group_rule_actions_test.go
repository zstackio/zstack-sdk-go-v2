// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityGroupRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySecurityGroupRule(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySecurityGroupRule result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Type, r.Protocol, r.State)
	}
	golog.Infof("======================================")
}

func TestPageSecurityGroupRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSecurityGroupRule(&queryParam)
	if err != nil {
		t.Errorf("TestPageSecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("PageSecurityGroupRule result: total=%d, returned=%d", total, len(result))
}

func TestGetSecurityGroupRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySecurityGroupRule(&queryParam)
	if err != nil {
		t.Errorf("TestGetSecurityGroupRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityGroupRule found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSecurityGroupRule(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("GetSecurityGroupRule result: %s", result.UUID)
}
