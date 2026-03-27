// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySdnController(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySdnController(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySdnController error: %v", err)
		return
	}
	golog.Infof("QuerySdnController result count: %d", len(result))
}

