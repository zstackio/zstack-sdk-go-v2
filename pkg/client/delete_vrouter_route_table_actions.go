// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteVRouterRouteTable deletes VRouterRouteTable
func (cli *ZSClient) DeleteVRouterRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{uuid}", uuid, string(deleteMode))
}
