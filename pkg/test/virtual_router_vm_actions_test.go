// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVirtualRouterVm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterVm result count: %d", len(result))
}
func TestGetVirtualRouterVm(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVirtualRouterVm(&queryParam)
	if err != nil {
		t.Errorf("TestGetVirtualRouterVm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VirtualRouterVm found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVirtualRouterVm(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("GetVirtualRouterVm result: %s", result.UUID)
}
