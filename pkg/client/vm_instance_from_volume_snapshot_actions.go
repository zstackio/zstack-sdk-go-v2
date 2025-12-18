// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmInstanceFromVolumeSnapshot 创建VmInstanceFromVolumeSnapshot
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshot(params param.CreateVmInstanceFromVolumeSnapshotParam) (*view.CreateVmInstanceFromVolumeSnapshotEventView, error) {
	resp := view.CreateVmInstanceFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/vm-instances/from/volume-snapshots/{volumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

