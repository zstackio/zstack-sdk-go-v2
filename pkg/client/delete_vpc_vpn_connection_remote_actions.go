// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcVpnConnectionRemote deletes VpcVpnConnectionRemote
func (cli *ZSClient) DeleteVpcVpnConnectionRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/{uuid}/remote", uuid, string(deleteMode))
}
