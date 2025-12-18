// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPrimaryStorageFromCluster 操作PrimaryStorageFromCluster
func (cli *ZSClient) DetachPrimaryStorageFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}", uuid, string(deleteMode))
}

