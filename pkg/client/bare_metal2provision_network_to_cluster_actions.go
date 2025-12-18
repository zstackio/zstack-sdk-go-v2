// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachBareMetal2ProvisionNetworkToCluster 操作BareMetal2ProvisionNetworkToCluster
func (cli *ZSClient) AttachBareMetal2ProvisionNetworkToCluster(params param.AttachBareMetal2ProvisionNetworkToClusterParam) (*view.AttachBareMetal2ProvisionNetworkToClusterEventView, error) {
	resp := view.AttachBareMetal2ProvisionNetworkToClusterEventView{}
	if err := cli.Post("v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

