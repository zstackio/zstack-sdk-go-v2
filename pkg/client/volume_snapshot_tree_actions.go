// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVolumeSnapshotTree queries VolumeSnapshotTree list
func (cli *ZSClient) QueryVolumeSnapshotTree(params *param.QueryParam) ([]view.VolumeSnapshotTreeInventoryView, error) {
	var resp []view.VolumeSnapshotTreeInventoryView
	return resp, cli.List("v1/volume-snapshots/trees", params, &resp)
}
