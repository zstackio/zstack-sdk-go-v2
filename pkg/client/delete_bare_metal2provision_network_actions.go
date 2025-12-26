// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBareMetal2ProvisionNetwork deletes BareMetal2ProvisionNetwork
func (cli *ZSClient) DeleteBareMetal2ProvisionNetwork(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/provision-networks/{uuid}", uuid, string(deleteMode))
}
