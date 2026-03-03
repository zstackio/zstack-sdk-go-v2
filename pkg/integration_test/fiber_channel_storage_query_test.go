// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFiberChannelStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryFiberChannelStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFiberChannelStorage error: %v", err)
		return
	}
	golog.Infof("QueryFiberChannelStorage result count: %d", len(result))
}

