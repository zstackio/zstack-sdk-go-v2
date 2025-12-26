// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckFirewallRuleConfigFile operates on CheckFirewallRuleConfigFile
func (cli *ZSClient) CheckFirewallRuleConfigFile(params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
