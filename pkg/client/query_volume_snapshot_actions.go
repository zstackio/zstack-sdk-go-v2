// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVolumeSnapshot queries VolumeSnapshot list
func (cli *ZSClient) QueryVolumeSnapshot(params *param.QueryParam) ([]view.VolumeSnapshotInventoryView, error) {
	var resp []view.VolumeSnapshotInventoryView
	return resp, cli.List("v1/volume-snapshots", params, &resp)
}
