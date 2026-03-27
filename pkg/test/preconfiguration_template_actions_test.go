// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPreconfigurationTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPreconfigurationTemplate(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQueryPreconfigurationTemplate error: %v", err)
		return
	}
	golog.Infof("QueryPreconfigurationTemplate result count: %d", len(result))
}

