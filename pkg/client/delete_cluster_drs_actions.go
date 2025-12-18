// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteClusterDRS deletes ClusterDRS
func (cli *ZSClient) DeleteClusterDRS(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/drs/{uuid}", uuid, string(deleteMode))
}
