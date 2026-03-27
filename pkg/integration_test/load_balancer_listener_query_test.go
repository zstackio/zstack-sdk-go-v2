// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancerListener(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancerListener result count: %d", len(result))
}

