// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteRuleSetVRouterRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteRuleSetVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteRuleSetVRouterRef error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteRuleSetVRouterRef result count: %d", len(result))
}
func TestGetPolicyRouteRuleSetVRouterRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteRuleSetVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetPolicyRouteRuleSetVRouterRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteRuleSetVRouterRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPolicyRouteRuleSetVRouterRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPolicyRouteRuleSetVRouterRef error: %v", err)
		return
	}
	golog.Infof("GetPolicyRouteRuleSetVRouterRef result: %s", result.UUID)
}
