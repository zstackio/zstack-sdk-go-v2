// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolumeFromVolumeSnapshot creates DataVolumeFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeFromVolumeSnapshot(params param.CreateDataVolumeFromVolumeSnapshotParam) (*view.CreateDataVolumeFromVolumeSnapshotEventView, error) {
	resp := view.CreateDataVolumeFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/volumes/data/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
