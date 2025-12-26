// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmInstanceFromVolumeSnapshot creates VmInstanceFromVolumeSnapshot
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshot(params param.CreateVmInstanceFromVolumeSnapshotParam) (*view.CreateVmInstanceFromVolumeSnapshotEventView, error) {
	resp := view.CreateVmInstanceFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/vm-instances/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
