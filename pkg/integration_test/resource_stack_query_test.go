// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryResourceStack(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryResourceStack(&queryParam)
	if err != nil {
		t.Errorf("TestQueryResourceStack error: %v", err)
		return
	}
	golog.Infof("QueryResourceStack result count: %d", len(result))
}

