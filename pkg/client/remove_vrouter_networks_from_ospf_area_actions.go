// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveVRouterNetworksFromOspfArea 操作RemoveVRouterNetworksFromOspfArea
func (cli *ZSClient) RemoveVRouterNetworksFromOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea/networks", uuid, string(deleteMode))
}

