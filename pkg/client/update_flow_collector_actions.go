// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateFlowCollector updates FlowCollector
func (cli *ZSClient) UpdateFlowCollector(uuid string, params param.UpdateFlowCollectorParam) (*view.CreateFlowCollectorEventView, error) {
	resp := view.CreateFlowCollectorEventView{}
	if err := cli.Put("v1/flowmeters/collectors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
