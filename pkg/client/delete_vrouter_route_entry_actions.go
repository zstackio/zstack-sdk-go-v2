// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVRouterRouteEntry deletes VRouterRouteEntry
func (cli *ZSClient) DeleteVRouterRouteEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{routeTableUuid}/route-entries/{uuid}", uuid, string(deleteMode))
}
