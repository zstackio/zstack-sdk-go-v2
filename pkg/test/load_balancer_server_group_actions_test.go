// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancerServerGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLoadBalancerServerGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancerServerGroup error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancerServerGroup result count: %d", len(result))
}

