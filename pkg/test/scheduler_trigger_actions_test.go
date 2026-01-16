// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySchedulerTrigger result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.SchedulerType)
	}
	golog.Infof("======================================")
}

func TestPageSchedulerTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestPageSchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("PageSchedulerTrigger result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.SchedulerType)
	}
}

func TestGetSchedulerTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySchedulerTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestGetSchedulerTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerTrigger found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSchedulerTrigger(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSchedulerTrigger error: %v", err)
		return
	}
	golog.Infof("GetSchedulerTrigger result: %s, Name: %s", result.UUID, result.Name)
}
