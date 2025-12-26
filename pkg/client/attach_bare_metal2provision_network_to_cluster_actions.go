// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachBareMetal2ProvisionNetworkToCluster operates on BareMetal2ProvisionNetworkToCluster
func (cli *ZSClient) AttachBareMetal2ProvisionNetworkToCluster(params param.AttachBareMetal2ProvisionNetworkToClusterParam) (*view.AttachBareMetal2ProvisionNetworkToClusterEventView, error) {
	resp := view.AttachBareMetal2ProvisionNetworkToClusterEventView{}
	if err := cli.Post("v1/baremetal2/clusters/{clusterUuid}/provision-networks/{networkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
