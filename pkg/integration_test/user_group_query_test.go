// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUserGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryUserGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUserGroup error: %v", err)
		return
	}
	golog.Infof("QueryUserGroup result count: %d", len(result))
}

