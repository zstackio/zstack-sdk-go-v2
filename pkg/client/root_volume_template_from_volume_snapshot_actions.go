// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateRootVolumeTemplateFromVolumeSnapshot 创建RootVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeSnapshot(params param.CreateRootVolumeTemplateFromVolumeSnapshotParam) (*view.CreateRootVolumeTemplateFromVolumeSnapshotEventView, error) {
	resp := view.CreateRootVolumeTemplateFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/images/root-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

