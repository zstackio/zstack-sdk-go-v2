// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVolumeSnapshot updates VolumeSnapshot
func (cli *ZSClient) UpdateVolumeSnapshot(uuid string, params param.UpdateVolumeSnapshotParam) (*view.UpdateVolumeSnapshotEventView, error) {
	resp := view.UpdateVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
