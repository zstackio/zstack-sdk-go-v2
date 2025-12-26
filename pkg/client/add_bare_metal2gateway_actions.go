// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddBareMetal2Gateway adds BareMetal2Gateway
func (cli *ZSClient) AddBareMetal2Gateway(params param.AddBareMetal2GatewayParam) (*view.AddHostEventView, error) {
	resp := view.AddHostEventView{}
	if err := cli.Post("v1/baremetal2/gateways", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
