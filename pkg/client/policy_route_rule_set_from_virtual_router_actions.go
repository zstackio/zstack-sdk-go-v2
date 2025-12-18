// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPolicyRouteRuleSetFromVirtualRouter 获取PolicyRouteRuleSetFromVirtualRouter详情
func (cli *ZSClient) GetPolicyRouteRuleSetFromVirtualRouter(uuid string) (*view.GetPolicyRouteRuleSetFromVirtualRouterView, error) {
	var resp view.GetPolicyRouteRuleSetFromVirtualRouterView
	if err := cli.Get("v1/policy-routes/rulesets/virtualrouter/{vmInstanceUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

