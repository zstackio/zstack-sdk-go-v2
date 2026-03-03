// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlockGroupPrimaryStorageHostRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySharedBlockGroupPrimaryStorageHostRef(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlockGroupPrimaryStorageHostRef error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlockGroupPrimaryStorageHostRef result count: %d", len(result))
}

