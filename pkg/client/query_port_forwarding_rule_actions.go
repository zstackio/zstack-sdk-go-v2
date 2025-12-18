// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPortForwardingRule queries PortForwardingRule list
func (cli *ZSClient) QueryPortForwardingRule(params param.QueryParam) ([]view.PortForwardingRuleInventoryView, error) {
	var resp []view.PortForwardingRuleInventoryView
	return resp, cli.List("v1/port-forwarding", &params, &resp)
}
