// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPhysicalSwitch(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPhysicalSwitch(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPhysicalSwitch error: %v", err)
		return
	}
	golog.Infof("QueryPhysicalSwitch result count: %d", len(result))
}
func TestGetPhysicalSwitch(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPhysicalSwitch(&queryParam)
	if err != nil {
		t.Errorf("TestGetPhysicalSwitch Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PhysicalSwitch found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPhysicalSwitch(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPhysicalSwitch error: %v", err)
		return
	}
	golog.Infof("GetPhysicalSwitch result: %s", result.UUID)
}
