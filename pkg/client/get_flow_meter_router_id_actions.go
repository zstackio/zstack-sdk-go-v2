// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetFlowMeterRouterId gets FlowMeterRouterId by uuid
func (cli *ZSClient) GetFlowMeterRouterId(uuid string) (*view.GetFlowMeterRouterIdView, error) {
	var resp view.GetFlowMeterRouterIdView
	if err := cli.Get("v1/flowmeters/{vRouterUuid}/routerid", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
