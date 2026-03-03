// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAppBuildSystem(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAppBuildSystem(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAppBuildSystem error: %v", err)
		return
	}
	golog.Infof("QueryAppBuildSystem result count: %d", len(result))
}

