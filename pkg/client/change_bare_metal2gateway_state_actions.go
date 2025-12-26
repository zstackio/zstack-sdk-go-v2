// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeBareMetal2GatewayState changes BareMetal2GatewayState
func (cli *ZSClient) ChangeBareMetal2GatewayState(uuid string, params param.ChangeBareMetal2GatewayStateParam) (*view.ChangeBareMetal2GatewayStateEventView, error) {
	resp := view.ChangeBareMetal2GatewayStateEventView{}
	if err := cli.Put("v1/baremetal2/gateways/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
