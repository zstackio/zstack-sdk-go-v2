// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcVpnConnectionLocal 删除VpcVpnConnectionLocal
func (cli *ZSClient) DeleteVpcVpnConnectionLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/{uuid}", uuid, string(deleteMode))
}

