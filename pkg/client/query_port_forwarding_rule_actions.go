// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPortForwardingRule queries PortForwardingRule list
func (cli *ZSClient) QueryPortForwardingRule(params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, error) {
	var resp []view.PortForwardingRuleInventoryView
	return resp, cli.List("v1/port-forwarding", params, &resp)
}
