// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteTableVRouterRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPolicyRouteTableVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteTableVRouterRef error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteTableVRouterRef result count: %d", len(result))
}

