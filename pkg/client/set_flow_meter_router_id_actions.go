// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetFlowMeterRouterId operates on SetFlowMeterRouterId
func (cli *ZSClient) SetFlowMeterRouterId(params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.Post("v1/flowmeters/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
