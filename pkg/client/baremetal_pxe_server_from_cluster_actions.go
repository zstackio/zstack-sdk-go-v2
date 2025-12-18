// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachBaremetalPxeServerFromCluster 操作BaremetalPxeServerFromCluster
func (cli *ZSClient) DetachBaremetalPxeServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/pxeservers/{pxeServerUuid}", uuid, string(deleteMode))
}

