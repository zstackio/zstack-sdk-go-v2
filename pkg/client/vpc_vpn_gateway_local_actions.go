// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcVpnGatewayLocal 删除VpcVpnGatewayLocal
func (cli *ZSClient) DeleteVpcVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-gateway/{uuid}", uuid, string(deleteMode))
}

