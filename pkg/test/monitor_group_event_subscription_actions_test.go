// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupEventSubscription(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroupEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupEventSubscription error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupEventSubscription result count: %d", len(result))
}
func TestGetMonitorGroupEventSubscription(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroupEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestGetMonitorGroupEventSubscription Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroupEventSubscription found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMonitorGroupEventSubscription(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMonitorGroupEventSubscription error: %v", err)
		return
	}
	golog.Infof("GetMonitorGroupEventSubscription result: %s", result.UUID)
}
