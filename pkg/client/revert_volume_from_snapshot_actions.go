// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RevertVolumeFromSnapshot 操作RevertVolumeFromSnapshot
func (cli *ZSClient) RevertVolumeFromSnapshot(uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.Put("v1/volume-snapshots/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

