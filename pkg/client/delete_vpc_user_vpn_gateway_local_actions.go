// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVpcUserVpnGatewayLocal deletes VpcUserVpnGatewayLocal
func (cli *ZSClient) DeleteVpcUserVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway/{uuid}", uuid, string(deleteMode))
}
