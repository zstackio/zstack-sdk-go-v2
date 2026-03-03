// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFiberChannelLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryFiberChannelLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFiberChannelLun error: %v", err)
		return
	}
	golog.Infof("QueryFiberChannelLun result count: %d", len(result))
}

