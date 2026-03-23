// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLocalStorageResourceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryLocalStorageResourceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLocalStorageResourceRef error: %v", err)
		return
	}
	golog.Infof("QueryLocalStorageResourceRef result count: %d", len(result))
}

