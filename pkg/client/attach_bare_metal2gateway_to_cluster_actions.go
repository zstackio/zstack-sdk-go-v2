// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachBareMetal2GatewayToCluster operates on BareMetal2GatewayToCluster
func (cli *ZSClient) AttachBareMetal2GatewayToCluster(params param.AttachBareMetal2GatewayToClusterParam) (*view.AttachBareMetal2GatewayToClusterEventView, error) {
	resp := view.AttachBareMetal2GatewayToClusterEventView{}
	if err := cli.Post("v1/baremetal2/clusters/{clusterUuid}/gateways/{gatewayUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
