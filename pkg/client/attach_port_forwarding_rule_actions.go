// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) AttachPortForwardingRule(params param.AttachPortForwardingRuleParam) (*view.AttachPortForwardingRuleEventView, error) {
	resp := view.AttachPortForwardingRuleEventView{}
	if err := cli.Post("v1/port-forwarding/{ruleUuid}/vm-instances/nics/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
