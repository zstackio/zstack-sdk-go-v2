// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVolumeSnapshotGroup 查询VolumeSnapshotGroup列表
func (cli *ZSClient) QueryVolumeSnapshotGroup(params param.QueryParam) ([]view.QueryVolumeSnapshotGroupView, error) {
	var resp []view.QueryVolumeSnapshotGroupView
	return resp, cli.List("v1/volume-snapshots/group", &params, &resp)
}

