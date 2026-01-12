// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterVRouterRouteTableRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVirtualRouterVRouterRouteTableRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVRouterRouteTableRef error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterVRouterRouteTableRef result count: %d", len(result))
}
func TestGetVirtualRouterVRouterRouteTableRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVirtualRouterVRouterRouteTableRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetVirtualRouterVRouterRouteTableRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VirtualRouterVRouterRouteTableRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVirtualRouterVRouterRouteTableRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVirtualRouterVRouterRouteTableRef error: %v", err)
		return
	}
	golog.Infof("GetVirtualRouterVRouterRouteTableRef result: %s", result.UUID)
}
