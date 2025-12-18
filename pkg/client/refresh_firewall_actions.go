// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshFirewall 操作RefreshFirewall
func (cli *ZSClient) RefreshFirewall(uuid string, params param.RefreshFirewallParam) (*view.RefreshFirewallEventView, error) {
	resp := view.RefreshFirewallEventView{}
	if err := cli.Put("v1/vpcfirewalls/refresh/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

