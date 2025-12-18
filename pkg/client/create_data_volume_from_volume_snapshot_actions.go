// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataVolumeFromVolumeSnapshot creates DataVolumeFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeFromVolumeSnapshot(params param.CreateDataVolumeFromVolumeSnapshotParam) (*view.CreateDataVolumeFromVolumeSnapshotEventView, error) {
	resp := view.CreateDataVolumeFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/volumes/data/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
