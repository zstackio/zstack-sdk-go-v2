// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBareMetal2Gateway updates BareMetal2Gateway
func (cli *ZSClient) UpdateBareMetal2Gateway(uuid string, params param.UpdateBareMetal2GatewayParam) (*view.UpdateBareMetal2GatewayEventView, error) {
	resp := view.UpdateBareMetal2GatewayEventView{}
	if err := cli.Put("v1/baremetal2/gateways/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
