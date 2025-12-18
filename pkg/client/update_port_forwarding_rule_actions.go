// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePortForwardingRule updates PortForwardingRule
func (cli *ZSClient) UpdatePortForwardingRule(uuid string, params param.UpdatePortForwardingRuleParam) (*view.UpdatePortForwardingRuleEventView, error) {
	resp := view.UpdatePortForwardingRuleEventView{}
	if err := cli.Put("v1/port-forwarding/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
