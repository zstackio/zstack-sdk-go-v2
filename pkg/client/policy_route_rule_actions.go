// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRule queries PolicyRouteRule list
func (cli *ZSClient) QueryPolicyRouteRule(params *param.QueryParam) ([]view.PolicyRouteRuleInventoryView, error) {
	var resp []view.PolicyRouteRuleInventoryView
	return resp, cli.List("v1/policy-routes/rules", params, &resp)
}
// CreatePolicyRouteRule creates PolicyRouteRule
func (cli *ZSClient) CreatePolicyRouteRule(params param.CreatePolicyRouteRuleParam) (*view.PolicyRouteRuleInventoryView, error) {
	var resp view.CreatePolicyRouteRuleEventView
	if err := cli.Post("v1/policy-routes/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePolicyRouteRule deletes PolicyRouteRule
func (cli *ZSClient) DeletePolicyRouteRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rules/{uuid}", uuid, string(deleteMode))
}
