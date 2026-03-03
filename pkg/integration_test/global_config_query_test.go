// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGlobalConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGlobalConfig error: %v", err)
		return
	}
	golog.Infof("QueryGlobalConfig result count: %d", len(result))
}

