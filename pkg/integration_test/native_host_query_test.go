// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNativeHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryNativeHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNativeHost error: %v", err)
		return
	}
	golog.Infof("QueryNativeHost result count: %d", len(result))
}

