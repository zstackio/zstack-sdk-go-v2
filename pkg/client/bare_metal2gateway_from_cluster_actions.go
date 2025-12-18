// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachBareMetal2GatewayFromCluster 操作BareMetal2GatewayFromCluster
func (cli *ZSClient) DetachBareMetal2GatewayFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/clusters/{clusterUuid}/gateways/{gatewayUuid}", uuid, string(deleteMode))
}

