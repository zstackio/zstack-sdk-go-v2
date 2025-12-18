// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachBareMetal2ProvisionNetworkFromCluster 操作BareMetal2ProvisionNetworkFromCluster
func (cli *ZSClient) DetachBareMetal2ProvisionNetworkFromCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", uuid, string(deleteMode))
}

