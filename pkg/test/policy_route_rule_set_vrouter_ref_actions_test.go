// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

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

