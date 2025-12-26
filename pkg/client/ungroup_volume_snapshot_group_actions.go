// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// UngroupVolumeSnapshotGroup operates on UngroupVolumeSnapshotGroup
func (cli *ZSClient) UngroupVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/ungroup/{uuid}", uuid, string(deleteMode))
}
