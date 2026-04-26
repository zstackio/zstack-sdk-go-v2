// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorageHostRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorageHostRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorageHostRef error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorageHostRef result count: %d", len(result))
}

