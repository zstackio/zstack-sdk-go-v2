// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachNvmeServerFromCluster 操作NvmeServerFromCluster
func (cli *ZSClient) DetachNvmeServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/storage-devices/nvme/servers/{uuid}", uuid, string(deleteMode))
}

