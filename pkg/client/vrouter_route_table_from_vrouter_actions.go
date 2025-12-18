// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachVRouterRouteTableFromVRouter 操作VRouterRouteTableFromVRouter
func (cli *ZSClient) DetachVRouterRouteTableFromVRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{routeTableUuid}/detach/{virtualRouterVmUuid}", uuid, string(deleteMode))
}

