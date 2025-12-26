// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmInstanceFromVolumeSnapshotGroup creates VmInstanceFromVolumeSnapshotGroup
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshotGroup(params param.CreateVmInstanceFromVolumeSnapshotGroupParam) (*view.CreateVmInstanceFromVolumeSnapshotGroupEventView, error) {
	resp := view.CreateVmInstanceFromVolumeSnapshotGroupEventView{}
	if err := cli.Post("v1/vm-instances/from/volume-snapshots/group/{volumeSnapshotGroupUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
