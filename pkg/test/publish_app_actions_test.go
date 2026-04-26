// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPublishApp(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPublishApp(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPublishApp error: %v", err)
		return
	}
	golog.Infof("QueryPublishApp result count: %d", len(result))
}

