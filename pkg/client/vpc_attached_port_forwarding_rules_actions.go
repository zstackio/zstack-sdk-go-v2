// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcAttachedPortForwardingRules 获取VpcAttachedPortForwardingRules详情
func (cli *ZSClient) GetVpcAttachedPortForwardingRules(uuid string) (*view.GetVpcAttachedPortForwardingRulesView, error) {
	var resp view.GetVpcAttachedPortForwardingRulesView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-pf", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

