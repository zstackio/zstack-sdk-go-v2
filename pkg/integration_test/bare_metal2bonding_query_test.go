// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Bonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBareMetal2Bonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Bonding error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Bonding result count: %d", len(result))
}

