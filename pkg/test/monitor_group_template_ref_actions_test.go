// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupTemplateRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroupTemplateRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupTemplateRef error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupTemplateRef result count: %d", len(result))
}
func TestGetMonitorGroupTemplateRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMonitorGroupTemplateRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetMonitorGroupTemplateRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MonitorGroupTemplateRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMonitorGroupTemplateRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMonitorGroupTemplateRef error: %v", err)
		return
	}
	golog.Infof("GetMonitorGroupTemplateRef result: %s", result.UUID)
}
