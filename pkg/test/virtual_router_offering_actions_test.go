// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVirtualRouterOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVirtualRouterOffering(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterOffering result count: %d", len(result))
}

