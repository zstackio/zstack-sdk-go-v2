// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVolumeSnapshotGroup creates VolumeSnapshotGroup
func (cli *ZSClient) CreateVolumeSnapshotGroup(params param.CreateVolumeSnapshotGroupParam) (*view.CreateVolumeSnapshotGroupEventView, error) {
	resp := view.CreateVolumeSnapshotGroupEventView{}
	if err := cli.Post("v1/volume-snapshots/group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
