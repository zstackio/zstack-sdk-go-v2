// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVolumeSnapshotGroup deletes VolumeSnapshotGroup
func (cli *ZSClient) DeleteVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/group/{uuid}", uuid, string(deleteMode))
}
