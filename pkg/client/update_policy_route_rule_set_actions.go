// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePolicyRouteRuleSet updates PolicyRouteRuleSet
func (cli *ZSClient) UpdatePolicyRouteRuleSet(uuid string, params param.UpdatePolicyRouteRuleSetParam) (*view.UpdatePolicyRouteRuleSetEventView, error) {
	resp := view.UpdatePolicyRouteRuleSetEventView{}
	if err := cli.Put("v1/policy-routes/ruleSets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
