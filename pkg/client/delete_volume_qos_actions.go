// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVolumeQos deletes VolumeQos
func (cli *ZSClient) DeleteVolumeQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/qos", uuid, string(deleteMode))
}
