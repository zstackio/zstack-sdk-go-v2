// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePolicyRouteRuleSet updates PolicyRouteRuleSet
func (cli *ZSClient) UpdatePolicyRouteRuleSet(uuid string, params param.UpdatePolicyRouteRuleSetParam) (*view.UpdatePolicyRouteRuleSetEventView, error) {
	resp := view.UpdatePolicyRouteRuleSetEventView{}
	if err := cli.Put("v1/policy-routes/ruleSets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
