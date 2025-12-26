// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachL2NetworkFromCluster operates on L2NetworkFromCluster
func (cli *ZSClient) DetachL2NetworkFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", uuid, string(deleteMode))
}
