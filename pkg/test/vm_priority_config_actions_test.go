// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmPriorityConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmPriorityConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmPriorityConfig error: %v", err)
		return
	}
	golog.Infof("QueryVmPriorityConfig result count: %d", len(result))
}
func TestGetVmPriorityConfig(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmPriorityConfig(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmPriorityConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmPriorityConfig found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmPriorityConfig(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmPriorityConfig error: %v", err)
		return
	}
	golog.Infof("GetVmPriorityConfig result: %s", result.UUID)
}
