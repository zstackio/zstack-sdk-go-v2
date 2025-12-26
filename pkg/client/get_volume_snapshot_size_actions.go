// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVolumeSnapshotSize gets VolumeSnapshotSize by uuid
func (cli *ZSClient) GetVolumeSnapshotSize(uuid string) (*view.GetVolumeSnapshotSizeEventView, error) {
	var resp view.GetVolumeSnapshotSizeEventView
	if err := cli.Get("v1/volume-snapshots/{uuid}/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
