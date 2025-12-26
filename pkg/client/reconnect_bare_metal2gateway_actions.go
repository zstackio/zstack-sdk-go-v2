// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectBareMetal2Gateway operates on ReconnectBareMetal2Gateway
func (cli *ZSClient) ReconnectBareMetal2Gateway(uuid string, params param.ReconnectBareMetal2GatewayParam) (*view.ReconnectBareMetal2GatewayEventView, error) {
	resp := view.ReconnectBareMetal2GatewayEventView{}
	if err := cli.Put("v1/baremetal2/gateways/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
