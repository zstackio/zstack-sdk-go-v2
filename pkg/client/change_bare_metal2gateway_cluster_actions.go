// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeBareMetal2GatewayCluster changes BareMetal2GatewayCluster
func (cli *ZSClient) ChangeBareMetal2GatewayCluster(uuid string, params param.ChangeBareMetal2GatewayClusterParam) (*view.ChangeBareMetal2GatewayClusterEventView, error) {
	resp := view.ChangeBareMetal2GatewayClusterEventView{}
	if err := cli.Put("v1/baremetal2/gateways/{gatewayUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
