// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccount(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAccount(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccount error: %v", err)
		return
	}
	golog.Infof("QueryAccount result count: %d", len(result))
}

