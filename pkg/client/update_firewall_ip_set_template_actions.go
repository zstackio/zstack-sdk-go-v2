// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateFirewallIpSetTemplate updates FirewallIpSetTemplate
func (cli *ZSClient) UpdateFirewallIpSetTemplate(uuid string, params param.UpdateFirewallIpSetTemplateParam) (*view.UpdateFirewallIpSetTemplateEventView, error) {
	resp := view.UpdateFirewallIpSetTemplateEventView{}
	if err := cli.Put("v1/vpcfirewalls/ipset/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
