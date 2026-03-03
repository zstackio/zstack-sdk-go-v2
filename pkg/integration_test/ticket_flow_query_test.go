// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTicketFlow(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryTicketFlow(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTicketFlow error: %v", err)
		return
	}
	golog.Infof("QueryTicketFlow result count: %d", len(result))
}

