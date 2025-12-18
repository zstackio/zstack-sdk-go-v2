// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetFlowMeterRouterId 操作SetFlowMeterRouterId
func (cli *ZSClient) SetFlowMeterRouterId(params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.Post("v1/flowmeters/{vRouterUuid}/routerid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

