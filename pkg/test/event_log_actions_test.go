// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEventLog(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEventLog(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEventLog error: %v", err)
		return
	}
	golog.Infof("QueryEventLog result count: %d", len(result))
}
func TestGetEventLog(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEventLog(&queryParam)
	if err != nil {
		t.Errorf("TestGetEventLog Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EventLog found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetEventLog(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetEventLog error: %v", err)
		return
	}
	golog.Infof("GetEventLog result: %s", result.UUID)
}
