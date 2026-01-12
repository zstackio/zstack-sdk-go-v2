// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteTableVRouterRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteTableVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteTableVRouterRef error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteTableVRouterRef result count: %d", len(result))
}
func TestGetPolicyRouteTableVRouterRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteTableVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetPolicyRouteTableVRouterRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteTableVRouterRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPolicyRouteTableVRouterRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPolicyRouteTableVRouterRef error: %v", err)
		return
	}
	golog.Infof("GetPolicyRouteTableVRouterRef result: %s", result.UUID)
}
