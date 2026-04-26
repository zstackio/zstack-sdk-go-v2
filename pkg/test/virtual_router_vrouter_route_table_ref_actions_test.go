// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterVRouterRouteTableRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVirtualRouterVRouterRouteTableRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVRouterRouteTableRef error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterVRouterRouteTableRef result count: %d", len(result))
}

