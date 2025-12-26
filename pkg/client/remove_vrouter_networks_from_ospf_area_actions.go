// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveVRouterNetworksFromOspfArea removes VRouterNetworksFromOspfArea
func (cli *ZSClient) RemoveVRouterNetworksFromOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea/networks", uuid, string(deleteMode))
}
