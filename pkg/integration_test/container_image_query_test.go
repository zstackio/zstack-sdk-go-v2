// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryContainerImage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryContainerImage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryContainerImage error: %v", err)
		return
	}
	golog.Infof("QueryContainerImage result count: %d", len(result))
}

