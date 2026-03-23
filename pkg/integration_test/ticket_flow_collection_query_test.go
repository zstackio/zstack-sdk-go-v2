// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTicketFlowCollection(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryTicketFlowCollection(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTicketFlowCollection error: %v", err)
		return
	}
	golog.Infof("QueryTicketFlowCollection result count: %d", len(result))
}

