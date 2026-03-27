// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
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

