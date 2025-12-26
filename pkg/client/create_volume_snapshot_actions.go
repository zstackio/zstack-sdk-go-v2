// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVolumeSnapshot creates VolumeSnapshot
func (cli *ZSClient) CreateVolumeSnapshot(params param.CreateVolumeSnapshotParam) (*view.CreateVolumeSnapshotEventView, error) {
	resp := view.CreateVolumeSnapshotEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
