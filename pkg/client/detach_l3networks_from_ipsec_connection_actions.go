// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachL3NetworksFromIPsecConnection operates on L3NetworksFromIPsecConnection
func (cli *ZSClient) DetachL3NetworksFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec/{uuid}/l3networks", uuid, string(deleteMode))
}
