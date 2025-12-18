// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveDnsFromVpcRouter 操作RemoveDnsFromVpcRouter
func (cli *ZSClient) RemoveDnsFromVpcRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/virtual-routers/{uuid}/dns", uuid, string(deleteMode))
}

