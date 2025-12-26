// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePolicyRouteRule creates PolicyRouteRule
func (cli *ZSClient) CreatePolicyRouteRule(params param.CreatePolicyRouteRuleParam) (*view.CreatePolicyRouteRuleEventView, error) {
	resp := view.CreatePolicyRouteRuleEventView{}
	if err := cli.Post("v1/policy-routes/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
