// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVpcIkeConfigLocal 删除VpcIkeConfigLocal
func (cli *ZSClient) DeleteVpcIkeConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ike/{uuid}", uuid, string(deleteMode))
}

