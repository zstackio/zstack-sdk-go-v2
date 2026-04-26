// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLoadBalancer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancer error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancer result count: %d", len(result))
}

