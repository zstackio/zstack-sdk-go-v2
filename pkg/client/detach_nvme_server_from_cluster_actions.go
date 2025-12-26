// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachNvmeServerFromCluster operates on NvmeServerFromCluster
func (cli *ZSClient) DetachNvmeServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", uuid, string(deleteMode))
}
