// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteRuleSetVRouterRef 查询PolicyRouteRuleSetVRouterRef列表
func (cli *ZSClient) QueryPolicyRouteRuleSetVRouterRef(params param.QueryParam) ([]view.QueryPolicyRouteRuleSetVRouterRefView, error) {
	var resp []view.QueryPolicyRouteRuleSetVRouterRefView
	return resp, cli.List("v1/policy-routes/rulesets/vrouters/refs", &params, &resp)
}

