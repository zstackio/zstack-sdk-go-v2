// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySecurityGroup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageSecurityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestPageSecurityGroup error: %v", err)
		return
	}
	golog.Infof("PageSecurityGroup result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
}

func TestGetSecurityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetSecurityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityGroup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSecurityGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSecurityGroup error: %v", err)
		return
	}
	golog.Infof("GetSecurityGroup result: %s, Name: %s", result.UUID, result.Name)
}
