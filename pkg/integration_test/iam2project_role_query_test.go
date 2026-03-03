// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectRole(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryIAM2ProjectRole(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectRole error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectRole result count: %d", len(result))
}

