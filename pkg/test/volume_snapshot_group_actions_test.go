// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshotGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVolumeSnapshotGroup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageVolumeSnapshotGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestPageVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("PageVolumeSnapshotGroup result: total=%d, returned=%d", total, len(result))
}

func TestGetVolumeSnapshotGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshotGroup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVolumeSnapshotGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("GetVolumeSnapshotGroup result: %s, Name: %s", result.UUID, result.Name)
}
