// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachIscsiServerFromCluster 操作IscsiServerFromCluster
func (cli *ZSClient) DetachIscsiServerFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", uuid, string(deleteMode))
}

