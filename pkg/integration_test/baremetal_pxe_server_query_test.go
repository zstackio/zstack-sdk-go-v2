// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBaremetalPxeServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBaremetalPxeServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalPxeServer error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalPxeServer result count: %d", len(result))
}

