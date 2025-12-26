// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachDataVolumeFromHost operates on DataVolumeFromHost
func (cli *ZSClient) DetachDataVolumeFromHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{volumeUuid}/hosts", uuid, string(deleteMode))
}
