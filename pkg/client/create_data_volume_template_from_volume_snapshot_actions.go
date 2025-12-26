// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolumeTemplateFromVolumeSnapshot creates DataVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeSnapshot(params param.CreateDataVolumeTemplateFromVolumeSnapshotParam) (*view.CreateDataVolumeTemplateFromVolumeSnapshotEventView, error) {
	resp := view.CreateDataVolumeTemplateFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/images/data-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
