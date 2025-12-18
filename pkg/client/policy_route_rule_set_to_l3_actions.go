// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPolicyRouteRuleSetToL3 操作PolicyRouteRuleSetToL3
func (cli *ZSClient) AttachPolicyRouteRuleSetToL3(params param.AttachPolicyRouteRuleSetToL3Param) (*view.AttachPolicyRouteRuleSetToL3EventView, error) {
	resp := view.AttachPolicyRouteRuleSetToL3EventView{}
	if err := cli.Post("v1/policy-routes/rulesets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

