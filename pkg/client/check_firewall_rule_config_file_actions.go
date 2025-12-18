// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckFirewallRuleConfigFile 操作CheckFirewallRuleConfigFile
func (cli *ZSClient) CheckFirewallRuleConfigFile(params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

