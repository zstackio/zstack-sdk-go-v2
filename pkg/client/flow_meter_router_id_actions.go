// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetFlowMeterRouterId 获取FlowMeterRouterId详情
func (cli *ZSClient) GetFlowMeterRouterId(uuid string) (*view.GetFlowMeterRouterIdView, error) {
	var resp view.GetFlowMeterRouterIdView
	if err := cli.Get("v1/flowmeters/{vRouterUuid}/routerid", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

