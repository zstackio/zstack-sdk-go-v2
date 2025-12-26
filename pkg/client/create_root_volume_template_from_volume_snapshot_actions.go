// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateRootVolumeTemplateFromVolumeSnapshot creates RootVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeSnapshot(params param.CreateRootVolumeTemplateFromVolumeSnapshotParam) (*view.CreateRootVolumeTemplateFromVolumeSnapshotEventView, error) {
	resp := view.CreateRootVolumeTemplateFromVolumeSnapshotEventView{}
	if err := cli.Post("v1/images/root-volume-templates/from/volume-snapshots/{snapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
