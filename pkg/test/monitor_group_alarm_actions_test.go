// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroupAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupAlarm error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupAlarm result count: %d", len(result))
}
func TestGetMonitorGroupAlarm(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroupAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestGetMonitorGroupAlarm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroupAlarm found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMonitorGroupAlarm(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMonitorGroupAlarm error: %v", err)
		return
	}
	golog.Infof("GetMonitorGroupAlarm result: %s", result.UUID)
}
