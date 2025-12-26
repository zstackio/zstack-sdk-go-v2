// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMdevDevice deletes MdevDevice
func (cli *ZSClient) DeleteMdevDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices/{mdevDeviceUuid}", uuid, string(deleteMode))
}
