// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVolumeSnapshot deletes VolumeSnapshot
func (cli *ZSClient) DeleteVolumeSnapshot(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/{uuid}", uuid, string(deleteMode))
}
