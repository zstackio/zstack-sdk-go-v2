// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVirtualBorderRouterLocal deletes VirtualBorderRouterLocal
func (cli *ZSClient) DeleteVirtualBorderRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/border-router/{uuid}", uuid, string(deleteMode))
}
