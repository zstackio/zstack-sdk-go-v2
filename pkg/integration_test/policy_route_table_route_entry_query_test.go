// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteTableRouteEntry(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPolicyRouteTableRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteTableRouteEntry error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteTableRouteEntry result count: %d", len(result))
}

