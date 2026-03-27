// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlockGroupPrimaryStorageHostRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedBlockGroupPrimaryStorageHostRef(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlockGroupPrimaryStorageHostRef error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlockGroupPrimaryStorageHostRef result count: %d", len(result))
}

