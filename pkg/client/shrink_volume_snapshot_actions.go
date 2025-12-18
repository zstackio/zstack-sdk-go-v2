// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ShrinkVolumeSnapshot 操作ShrinkVolumeSnapshot
func (cli *ZSClient) ShrinkVolumeSnapshot(uuid string, params param.ShrinkVolumeSnapshotParam) (*view.ShrinkVolumeSnapshotEventView, error) {
	resp := view.ShrinkVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/shrink/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

