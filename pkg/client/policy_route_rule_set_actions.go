// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteRuleSet 查询PolicyRouteRuleSet列表
func (cli *ZSClient) QueryPolicyRouteRuleSet(params param.QueryParam) ([]view.QueryPolicyRouteRuleSetView, error) {
	var resp []view.QueryPolicyRouteRuleSetView
	return resp, cli.List("v1/policy-routes/rulesets", &params, &resp)
}

