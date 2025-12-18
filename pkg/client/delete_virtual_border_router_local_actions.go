// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVirtualBorderRouterLocal deletes VirtualBorderRouterLocal
func (cli *ZSClient) DeleteVirtualBorderRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/border-router/{uuid}", uuid, string(deleteMode))
}
