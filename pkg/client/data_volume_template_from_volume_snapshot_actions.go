// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataVolumeTemplateFromVolumeSnapshot 创建DataVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeSnapshot(params param.CreateDataVolumeTemplateFromVolumeSnapshotParam) (*view.CreateDataVolumeTemplateFromVolumeSnapshotEventView, error) {
	resp := view.CreateDataVolumeTemplateFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/images/data-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

