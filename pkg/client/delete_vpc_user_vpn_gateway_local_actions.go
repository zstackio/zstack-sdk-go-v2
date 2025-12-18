// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcUserVpnGatewayLocal deletes VpcUserVpnGatewayLocal
func (cli *ZSClient) DeleteVpcUserVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway/{uuid}", uuid, string(deleteMode))
}
