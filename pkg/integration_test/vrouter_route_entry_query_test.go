// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVRouterRouteEntry(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVRouterRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVRouterRouteEntry error: %v", err)
		return
	}
	golog.Infof("QueryVRouterRouteEntry result count: %d", len(result))
}

