// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeTarget(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNvmeTarget(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeTarget error: %v", err)
		return
	}
	golog.Infof("QueryNvmeTarget result count: %d", len(result))
}

