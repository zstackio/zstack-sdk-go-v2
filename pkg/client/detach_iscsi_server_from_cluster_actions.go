// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachIscsiServerFromCluster operates on IscsiServerFromCluster
func (cli *ZSClient) DetachIscsiServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", uuid, string(deleteMode))
}
