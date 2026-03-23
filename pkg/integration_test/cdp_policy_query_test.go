// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCdpPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCdpPolicy error: %v", err)
		return
	}
	golog.Infof("QueryCdpPolicy result count: %d", len(result))
}

