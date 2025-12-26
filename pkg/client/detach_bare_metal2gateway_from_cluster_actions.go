// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachBareMetal2GatewayFromCluster operates on BareMetal2GatewayFromCluster
func (cli *ZSClient) DetachBareMetal2GatewayFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/clusters/{clusterUuid}/gateways/{gatewayUuid}", uuid, string(deleteMode))
}
