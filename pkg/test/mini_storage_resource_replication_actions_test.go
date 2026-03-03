// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorageResourceReplication(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorageResourceReplication(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorageResourceReplication error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorageResourceReplication result count: %d", len(result))
}

