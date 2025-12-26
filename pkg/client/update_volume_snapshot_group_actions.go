// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVolumeSnapshotGroup updates VolumeSnapshotGroup
func (cli *ZSClient) UpdateVolumeSnapshotGroup(uuid string, params param.UpdateVolumeSnapshotGroupParam) (*view.UpdateVolumeSnapshotGroupEventView, error) {
	resp := view.UpdateVolumeSnapshotGroupEventView{}
	if err := cli.Put("v1/volume-snapshots/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
