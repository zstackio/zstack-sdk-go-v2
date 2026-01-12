// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryActiveAlarm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryActiveAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryActiveAlarm error: %v", err)
		return
	}
	golog.Infof("QueryActiveAlarm result count: %d", len(result))
}
func TestGetActiveAlarm(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryActiveAlarm(&queryParam)
	if err != nil {
		t.Errorf("TestGetActiveAlarm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ActiveAlarm found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetActiveAlarm(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetActiveAlarm error: %v", err)
		return
	}
	golog.Infof("GetActiveAlarm result: %s", result.UUID)
}
