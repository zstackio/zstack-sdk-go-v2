// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVolumesSnapshot creates VolumesSnapshot
func (cli *ZSClient) CreateVolumesSnapshot(params param.CreateVolumesSnapshotParam) (*view.CreateVolumesSnapshotEventView, error) {
	resp := view.CreateVolumesSnapshotEventView{}
	if err := cli.Post("v1/volumes/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
