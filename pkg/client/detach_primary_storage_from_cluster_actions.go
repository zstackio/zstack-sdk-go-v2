// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPrimaryStorageFromCluster operates on PrimaryStorageFromCluster
func (cli *ZSClient) DetachPrimaryStorageFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}", uuid, string(deleteMode))
}
