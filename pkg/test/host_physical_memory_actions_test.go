// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostPhysicalMemory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostPhysicalMemory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostPhysicalMemory error: %v", err)
		return
	}
	golog.Infof("QueryHostPhysicalMemory result count: %d", len(result))
}
func TestGetHostPhysicalMemory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostPhysicalMemory(&queryParam)
	if err != nil {
		t.Errorf("TestGetHostPhysicalMemory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostPhysicalMemory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetHostPhysicalMemory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHostPhysicalMemory error: %v", err)
		return
	}
	golog.Infof("GetHostPhysicalMemory result: %s", result.UUID)
}
