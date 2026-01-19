// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAlarm error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryAlarm result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestPageAlarm error: %v", err)
		return
	}
	golog.Infof("PageAlarm result: total=%d, returned=%d", total, len(result))
}

func TestGetAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestGetAlarm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Alarm found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetAlarm(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAlarm error: %v", err)
		return
	}
	golog.Infof("GetAlarm result: %s, Name: %s", result.UUID, result.Name)
}
