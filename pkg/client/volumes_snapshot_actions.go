// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVolumesSnapshot 创建VolumesSnapshot
func (cli *ZSClient) CreateVolumesSnapshot(params param.CreateVolumesSnapshotParam) (*view.CreateVolumesSnapshotEventView, error) {
	resp := view.CreateVolumesSnapshotEventView{}
	if err := cli.Post("v1/volumes/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

