// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVpcIpSecConfigLocal deletes VpcIpSecConfigLocal
func (cli *ZSClient) DeleteVpcIpSecConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ipsec/{uuid}", uuid, string(deleteMode))
}
