// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPolicyRouteRuleSetToL3 operates on PolicyRouteRuleSetToL3
func (cli *ZSClient) AttachPolicyRouteRuleSetToL3(params param.AttachPolicyRouteRuleSetToL3Param) (*view.AttachPolicyRouteRuleSetToL3EventView, error) {
	resp := view.AttachPolicyRouteRuleSetToL3EventView{}
	if err := cli.Post("v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
