// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVolumeSnapshotTree 查询VolumeSnapshotTree列表
func (cli *ZSClient) QueryVolumeSnapshotTree(params param.QueryParam) ([]view.QueryVolumeSnapshotTreeView, error) {
	var resp []view.QueryVolumeSnapshotTreeView
	return resp, cli.List("v1/volume-snapshots/trees", &params, &resp)
}

