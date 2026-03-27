// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEmailEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSEmailEndpoint(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailEndpoint result count: %d", len(result))
}

