// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualIDAttribute(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryIAM2VirtualIDAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualIDAttribute error: %v", err)
		return
	}
	golog.Infof("QueryIAM2VirtualIDAttribute result count: %d", len(result))
}

