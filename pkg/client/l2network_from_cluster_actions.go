// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachL2NetworkFromCluster 操作L2NetworkFromCluster
func (cli *ZSClient) DetachL2NetworkFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", uuid, string(deleteMode))
}

