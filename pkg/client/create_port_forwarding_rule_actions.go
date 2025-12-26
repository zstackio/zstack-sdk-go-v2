// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePortForwardingRule creates PortForwardingRule
func (cli *ZSClient) CreatePortForwardingRule(params param.CreatePortForwardingRuleParam) (*view.CreatePortForwardingRuleEventView, error) {
	resp := view.CreatePortForwardingRuleEventView{}
	if err := cli.Post("v1/port-forwarding", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
