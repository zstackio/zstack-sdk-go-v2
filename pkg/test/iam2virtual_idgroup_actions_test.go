// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualIDGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2VirtualIDGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualIDGroup error: %v", err)
		return
	}
	golog.Infof("QueryIAM2VirtualIDGroup result count: %d", len(result))
}

