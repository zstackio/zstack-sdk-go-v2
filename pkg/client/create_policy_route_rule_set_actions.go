// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePolicyRouteRuleSet creates PolicyRouteRuleSet
func (cli *ZSClient) CreatePolicyRouteRuleSet(params param.CreatePolicyRouteRuleSetParam) (*view.CreatePolicyRouteRuleSetEventView, error) {
	resp := view.CreatePolicyRouteRuleSetEventView{}
	if err := cli.Post("v1/policy-routes/rulesets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
