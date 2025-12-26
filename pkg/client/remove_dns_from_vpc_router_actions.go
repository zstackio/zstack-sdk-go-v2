// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveDnsFromVpcRouter removes DnsFromVpcRouter
func (cli *ZSClient) RemoveDnsFromVpcRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/virtual-routers/{uuid}/dns", uuid, string(deleteMode))
}
