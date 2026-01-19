// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryAutoScalingGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingGroup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryAutoScalingGroup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageAutoScalingGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageAutoScalingGroup(&queryParam)
	if err != nil {
		t.Errorf("TestPageAutoScalingGroup error: %v", err)
		return
	}
	golog.Infof("PageAutoScalingGroup result: total=%d, returned=%d", total, len(result))
}

func TestGetAutoScalingGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryAutoScalingGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetAutoScalingGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetAutoScalingGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAutoScalingGroup error: %v", err)
		return
	}
	golog.Infof("GetAutoScalingGroup result: %s, Name: %s", result.UUID, result.Name)
}
