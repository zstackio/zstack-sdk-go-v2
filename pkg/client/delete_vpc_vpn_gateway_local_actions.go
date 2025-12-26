// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVpcVpnGatewayLocal deletes VpcVpnGatewayLocal
func (cli *ZSClient) DeleteVpcVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-gateway/{uuid}", uuid, string(deleteMode))
}
