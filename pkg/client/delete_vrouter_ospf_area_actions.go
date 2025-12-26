// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVRouterOspfArea deletes VRouterOspfArea
func (cli *ZSClient) DeleteVRouterOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea/{uuid}", uuid, string(deleteMode))
}
