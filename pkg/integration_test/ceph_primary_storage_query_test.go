// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephPrimaryStorage result count: %d", len(result))
}

