// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVolumeSnapshotTree queries VolumeSnapshotTree list
func (cli *ZSClient) QueryVolumeSnapshotTree(params *param.QueryParam) ([]view.VolumeSnapshotTreeInventoryView, error) {
	var resp []view.VolumeSnapshotTreeInventoryView
	return resp, cli.List("v1/volume-snapshots/trees", params, &resp)
}

// PageVolumeSnapshotTree Pagination
func (cli *ZSClient) PageVolumeSnapshotTree(params *param.QueryParam) ([]view.VolumeSnapshotTreeInventoryView, int, error) {
	var volumeSnapshotTrees []view.VolumeSnapshotTreeInventoryView
	total, err := cli.Page("v1/volume-snapshots/trees", params, &volumeSnapshotTrees)
	return volumeSnapshotTrees, total, err
}
