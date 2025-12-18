// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BatchDeleteVolumeSnapshot operates on BatchDeleteVolumeSnapshot
func (cli *ZSClient) BatchDeleteVolumeSnapshot(uuid string, params param.BatchDeleteVolumeSnapshotParam) (*view.BatchDeleteVolumeSnapshotEventView, error) {
	resp := view.BatchDeleteVolumeSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/batch-delete", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
