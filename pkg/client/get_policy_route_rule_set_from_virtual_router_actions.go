// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPolicyRouteRuleSetFromVirtualRouter gets PolicyRouteRuleSetFromVirtualRouter by uuid
func (cli *ZSClient) GetPolicyRouteRuleSetFromVirtualRouter(uuid string) (*view.GetPolicyRouteRuleSetFromVirtualRouterView, error) {
	var resp view.GetPolicyRouteRuleSetFromVirtualRouterView
	if err := cli.Get("v1/policy-routes/rulesets/virtualrouter/{vmInstanceUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
