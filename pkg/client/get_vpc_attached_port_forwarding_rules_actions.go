// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcAttachedPortForwardingRules gets VpcAttachedPortForwardingRules by uuid
func (cli *ZSClient) GetVpcAttachedPortForwardingRules(uuid string) (*view.GetVpcAttachedPortForwardingRulesView, error) {
	var resp view.GetVpcAttachedPortForwardingRulesView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-pf", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
