// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAffinityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAffinityGroup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryAffinityGroup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Policy)
	}
	golog.Infof("======================================")
}

func TestPageAffinityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestPageAffinityGroup error: %v", err)
		return
	}
	golog.Infof("PageAffinityGroup result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Policy)
	}
}

func TestGetAffinityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetAffinityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AffinityGroup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetAffinityGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAffinityGroup error: %v", err)
		return
	}
	golog.Infof("GetAffinityGroup result: %s, Name: %s", result.UUID, result.Name)
}
