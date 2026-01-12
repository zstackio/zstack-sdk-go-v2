// Copyright (c) ZStack.io, Inc.

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
func TestGetVolumeSnapshotTree(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshotTree(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotTree Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshotTree found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVolumeSnapshotTree(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotTree error: %v", err)
		return
	}
	golog.Infof("GetVolumeSnapshotTree result: %s", result.UUID)
}
