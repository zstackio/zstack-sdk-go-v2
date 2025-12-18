// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteRuleSetL3Ref 查询PolicyRouteRuleSetL3Ref列表
func (cli *ZSClient) QueryPolicyRouteRuleSetL3Ref(params param.QueryParam) ([]view.QueryPolicyRouteRuleSetL3RefView, error) {
	var resp []view.QueryPolicyRouteRuleSetL3RefView
	return resp, cli.List("v1/policy-routes/rulesets/l3networdks/refs", &params, &resp)
}

