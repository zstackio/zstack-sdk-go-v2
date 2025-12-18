// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFlowCollector updates FlowCollector
func (cli *ZSClient) UpdateFlowCollector(uuid string, params param.UpdateFlowCollectorParam) (*view.CreateFlowCollectorEventView, error) {
	resp := view.CreateFlowCollectorEventView{}
	if err := cli.Put("v1/flowmeters/collectors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
