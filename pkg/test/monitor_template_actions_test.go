// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryMonitorTemplate result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.MonitorGroupTemplateRefs)
	}
	golog.Infof("======================================")
}

func TestPageMonitorTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestPageMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("PageMonitorTemplate result: total=%d, returned=%d", total, len(result))
}

func TestGetMonitorTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryMonitorTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestGetMonitorTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorTemplate found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetMonitorTemplate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMonitorTemplate error: %v", err)
		return
	}
	golog.Infof("GetMonitorTemplate result: %s, Name: %s", result.UUID, result.Name)
}
