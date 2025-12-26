// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFlowCollector queries FlowCollector list
func (cli *ZSClient) QueryFlowCollector(params *param.QueryParam) ([]view.FlowCollectorInventoryView, error) {
	var resp []view.FlowCollectorInventoryView
	return resp, cli.List("v1/flowmeters/collectors", params, &resp)
}
