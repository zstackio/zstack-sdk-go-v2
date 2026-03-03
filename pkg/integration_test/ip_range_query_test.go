// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIpRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIpRange error: %v", err)
		return
	}
	golog.Infof("QueryIpRange result count: %d", len(result))
}

