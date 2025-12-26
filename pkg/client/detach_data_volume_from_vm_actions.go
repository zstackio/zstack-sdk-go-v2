// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachDataVolumeFromVm operates on DataVolumeFromVm
func (cli *ZSClient) DetachDataVolumeFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}/vm-instances", uuid, string(deleteMode))
}
