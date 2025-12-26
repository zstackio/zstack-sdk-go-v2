// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBareMetal2Chassis deletes BareMetal2Chassis
func (cli *ZSClient) DeleteBareMetal2Chassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/chassis/{uuid}", uuid, string(deleteMode))
}
