// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshFirewall operates on RefreshFirewall
func (cli *ZSClient) RefreshFirewall(uuid string, params param.RefreshFirewallParam) (*view.RefreshFirewallEventView, error) {
	resp := view.RefreshFirewallEventView{}
	if err := cli.Put("v1/vpcfirewalls/refresh/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
