// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVirtualRouterVm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterVm result count: %d", len(result))
}
