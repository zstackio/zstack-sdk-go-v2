// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFirewallIpSetTemplate updates FirewallIpSetTemplate
func (cli *ZSClient) UpdateFirewallIpSetTemplate(uuid string, params param.UpdateFirewallIpSetTemplateParam) (*view.UpdateFirewallIpSetTemplateEventView, error) {
	resp := view.UpdateFirewallIpSetTemplateEventView{}
	if err := cli.Put("v1/vpcfirewalls/ipset/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
