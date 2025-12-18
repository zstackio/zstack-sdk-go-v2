// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFirewallIpSetTemplate creates FirewallIpSetTemplate
func (cli *ZSClient) CreateFirewallIpSetTemplate(params param.CreateFirewallIpSetTemplateParam) (*view.CreateFirewallIpSetTemplateEventView, error) {
	resp := view.CreateFirewallIpSetTemplateEventView{}
	if err := cli.Post("v1/vpcfirewalls/ipset/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
