// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBaremetalChassis deletes BaremetalChassis
func (cli *ZSClient) DeleteBaremetalChassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/chassis/{uuid}", uuid, string(deleteMode))
}
