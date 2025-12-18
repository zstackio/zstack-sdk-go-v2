// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteCluster deletes Cluster
func (cli *ZSClient) DeleteCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{uuid}", uuid, string(deleteMode))
}
