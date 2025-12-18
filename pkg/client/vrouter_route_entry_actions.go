// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVRouterRouteEntry 删除VRouterRouteEntry
func (cli *ZSClient) DeleteVRouterRouteEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{routeTableUuid}/route-entries/{uuid}", uuid, string(deleteMode))
}

