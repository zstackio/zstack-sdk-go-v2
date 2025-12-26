// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFlowCollector creates FlowCollector
func (cli *ZSClient) CreateFlowCollector(params param.CreateFlowCollectorParam) (*view.CreateFlowCollectorEventView, error) {
	resp := view.CreateFlowCollectorEventView{}
	if err := cli.Post("v1/flowmeters/{flowMeterUuid}/collectors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
