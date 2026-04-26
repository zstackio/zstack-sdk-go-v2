// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshotTree(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVolumeSnapshotTree(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshotTree error: %v", err)
		return
	}
	golog.Infof("QueryVolumeSnapshotTree result count: %d", len(result))
}

