// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshotGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("QueryVolumeSnapshotGroup result count: %d", len(result))
}

