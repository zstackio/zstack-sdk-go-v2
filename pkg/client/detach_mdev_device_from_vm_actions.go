// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachMdevDeviceFromVm operates on MdevDeviceFromVm
func (cli *ZSClient) DetachMdevDeviceFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices/{mdevDeviceUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}
