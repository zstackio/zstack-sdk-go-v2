// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcIpSecConfigLocal 删除VpcIpSecConfigLocal
func (cli *ZSClient) DeleteVpcIpSecConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ipsec/{uuid}", uuid, string(deleteMode))
}

