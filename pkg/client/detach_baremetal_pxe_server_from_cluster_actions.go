// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachBaremetalPxeServerFromCluster operates on BaremetalPxeServerFromCluster
func (cli *ZSClient) DetachBaremetalPxeServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", uuid, string(deleteMode))
}
