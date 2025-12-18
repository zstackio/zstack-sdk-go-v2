// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteRule queries PolicyRouteRule list
func (cli *ZSClient) QueryPolicyRouteRule(params param.QueryParam) ([]view.PolicyRouteRuleInventoryView, error) {
	var resp []view.PolicyRouteRuleInventoryView
	return resp, cli.List("v1/policy-routes/rules", &params, &resp)
}
