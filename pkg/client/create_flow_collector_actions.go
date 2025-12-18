// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFlowCollector creates FlowCollector
func (cli *ZSClient) CreateFlowCollector(params param.CreateFlowCollectorParam) (*view.CreateFlowCollectorEventView, error) {
	resp := view.CreateFlowCollectorEventView{}
	if err := cli.Post("v1/flowmeters/{flowMeterUuid}/collectors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
