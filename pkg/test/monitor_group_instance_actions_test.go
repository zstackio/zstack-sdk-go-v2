// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupInstance error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupInstance result count: %d", len(result))
}
func TestGetMonitorGroupInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestGetMonitorGroupInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroupInstance found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMonitorGroupInstance(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMonitorGroupInstance error: %v", err)
		return
	}
	golog.Infof("GetMonitorGroupInstance result: %s", result.UUID)
}
