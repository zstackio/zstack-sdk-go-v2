// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectAttribute(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryIAM2ProjectAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectAttribute error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectAttribute result count: %d", len(result))
}

