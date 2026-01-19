// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshotTree(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVolumeSnapshotTree(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshotTree error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVolumeSnapshotTree result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.VolumeUuid)
	}
	golog.Infof("======================================")
}

func TestPageVolumeSnapshotTree(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVolumeSnapshotTree(&queryParam)
	if err != nil {
		t.Errorf("TestPageVolumeSnapshotTree error: %v", err)
		return
	}
	golog.Infof("PageVolumeSnapshotTree result: total=%d, returned=%d", total, len(result))
}

func TestGetVolumeSnapshotTree(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVolumeSnapshotTree(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotTree Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshotTree found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVolumeSnapshotTree(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotTree error: %v", err)
		return
	}
	golog.Infof("GetVolumeSnapshotTree result: %s", result.UUID)
}
