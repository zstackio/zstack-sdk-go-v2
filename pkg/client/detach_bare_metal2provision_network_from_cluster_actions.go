// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachBareMetal2ProvisionNetworkFromCluster operates on BareMetal2ProvisionNetworkFromCluster
func (cli *ZSClient) DetachBareMetal2ProvisionNetworkFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", uuid, string(deleteMode))
}
