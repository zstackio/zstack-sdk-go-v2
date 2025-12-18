// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFlowMeter updates FlowMeter
func (cli *ZSClient) UpdateFlowMeter(uuid string, params param.UpdateFlowMeterParam) (*view.UpdateFlowMeterEventView, error) {
	resp := view.UpdateFlowMeterEventView{}
	if err := cli.Put("v1/flowmeters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
