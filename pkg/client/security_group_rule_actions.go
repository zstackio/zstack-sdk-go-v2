// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySecurityGroupRule 查询SecurityGroupRule列表
func (cli *ZSClient) QuerySecurityGroupRule(params param.QueryParam) ([]view.QuerySecurityGroupRuleView, error) {
	var resp []view.QuerySecurityGroupRuleView
	return resp, cli.List("v1/security-groups/rules", &params, &resp)
}

