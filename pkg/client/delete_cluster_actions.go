// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteCluster deletes Cluster
func (cli *ZSClient) DeleteCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{uuid}", uuid, string(deleteMode))
}
