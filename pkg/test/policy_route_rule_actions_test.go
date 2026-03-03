// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteRule error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteRule result count: %d", len(result))
}

