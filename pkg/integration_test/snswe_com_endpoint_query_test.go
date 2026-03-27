// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSWeComEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSWeComEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSWeComEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSWeComEndpoint result count: %d", len(result))
}

