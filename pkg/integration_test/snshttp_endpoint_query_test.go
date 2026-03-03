// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSHttpEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSHttpEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSHttpEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSHttpEndpoint result count: %d", len(result))
}

