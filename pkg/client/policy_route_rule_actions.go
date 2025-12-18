// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteRule 查询PolicyRouteRule列表
func (cli *ZSClient) QueryPolicyRouteRule(params param.QueryParam) ([]view.QueryPolicyRouteRuleView, error) {
	var resp []view.QueryPolicyRouteRuleView
	return resp, cli.List("v1/policy-routes/rules", &params, &resp)
}

