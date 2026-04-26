// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDataset(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDataset(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDataset error: %v", err)
		return
	}
	golog.Infof("QueryDataset result count: %d", len(result))
}

