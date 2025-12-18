// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcUserVpnGatewayRemote deletes VpcUserVpnGatewayRemote
func (cli *ZSClient) DeleteVpcUserVpnGatewayRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway/{uuid}/remote", uuid, string(deleteMode))
}
