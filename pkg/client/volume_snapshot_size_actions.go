// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVolumeSnapshotSize 获取VolumeSnapshotSize详情
func (cli *ZSClient) GetVolumeSnapshotSize(uuid string) (*view.GetVolumeSnapshotSizeEventView, error) {
	var resp view.GetVolumeSnapshotSizeEventView
	if err := cli.Get("v1/volume-snapshots/{uuid}/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

