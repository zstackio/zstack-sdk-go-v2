// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRuleSetL3Ref queries PolicyRouteRuleSetL3Ref list
func (cli *ZSClient) QueryPolicyRouteRuleSetL3Ref(params *param.QueryParam) ([]view.PolicyRouteRuleSetL3RefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetL3RefInventoryView
	return resp, cli.List("v1/policy-routes/rulesets/l3networdks/refs", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRuleSetL3Ref(uuid string) (*view.PolicyRouteRuleSetL3RefInventoryView, error) {
	var resp view.PolicyRouteRuleSetL3RefInventoryView
	if err := cli.Get("v1/policy-routes/rulesets/l3networdks/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
