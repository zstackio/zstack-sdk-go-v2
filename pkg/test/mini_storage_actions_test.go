// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorage error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorage result count: %d", len(result))
}

