// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteClusterDRS deletes ClusterDRS
func (cli *ZSClient) DeleteClusterDRS(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/drs/{uuid}", uuid, string(deleteMode))
}
