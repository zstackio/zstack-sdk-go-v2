// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePortForwardingRule creates PortForwardingRule
func (cli *ZSClient) CreatePortForwardingRule(params param.CreatePortForwardingRuleParam) (*view.CreatePortForwardingRuleEventView, error) {
	resp := view.CreatePortForwardingRuleEventView{}
	if err := cli.Post("v1/port-forwarding", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
