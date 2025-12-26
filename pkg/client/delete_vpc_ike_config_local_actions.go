// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVpcIkeConfigLocal deletes VpcIkeConfigLocal
func (cli *ZSClient) DeleteVpcIkeConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ike/{uuid}", uuid, string(deleteMode))
}
