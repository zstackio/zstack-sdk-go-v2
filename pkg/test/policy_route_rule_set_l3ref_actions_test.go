// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteRuleSetL3Ref(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteRuleSetL3Ref(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteRuleSetL3Ref error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteRuleSetL3Ref result count: %d", len(result))
}
func TestGetPolicyRouteRuleSetL3Ref(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteRuleSetL3Ref(&queryParam)
	if err != nil {
		t.Errorf("TestGetPolicyRouteRuleSetL3Ref Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteRuleSetL3Ref found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPolicyRouteRuleSetL3Ref(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPolicyRouteRuleSetL3Ref error: %v", err)
		return
	}
	golog.Infof("GetPolicyRouteRuleSetL3Ref result: %s", result.UUID)
}
