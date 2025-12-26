// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RevertVolumeFromSnapshot operates on RevertVolumeFromSnapshot
func (cli *ZSClient) RevertVolumeFromSnapshot(uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
