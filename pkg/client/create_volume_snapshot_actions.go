// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVolumeSnapshot creates VolumeSnapshot
func (cli *ZSClient) CreateVolumeSnapshot(params param.CreateVolumeSnapshotParam) (*view.CreateVolumeSnapshotEventView, error) {
	resp := view.CreateVolumeSnapshotEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
