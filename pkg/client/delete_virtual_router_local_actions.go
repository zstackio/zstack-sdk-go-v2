// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVirtualRouterLocal deletes VirtualRouterLocal
func (cli *ZSClient) DeleteVirtualRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vrouter/{uuid}", uuid, string(deleteMode))
}
