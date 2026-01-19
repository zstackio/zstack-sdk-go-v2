// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserGroup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryUserGroup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageUserGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestPageUserGroup error: %v", err)
		return
	}
	golog.Infof("PageUserGroup result: total=%d, returned=%d", total, len(result))
}

func TestGetUserGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetUserGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No UserGroup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetUserGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetUserGroup error: %v", err)
		return
	}
	golog.Infof("GetUserGroup result: %s, Name: %s", result.UUID, result.Name)
}
