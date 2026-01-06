// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePolicyRouteRuleSet creates PolicyRouteRuleSet
func (cli *ZSClient) CreatePolicyRouteRuleSet(params param.CreatePolicyRouteRuleSetParam) (*view.PolicyRouteRuleSetInventoryView, error) {
	var resp view.CreatePolicyRouteRuleSetEventView
	if err := cli.Post("v1/policy-routes/rulesets", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePolicyRouteRuleSet updates PolicyRouteRuleSet
func (cli *ZSClient) UpdatePolicyRouteRuleSet(uuid string, params param.UpdatePolicyRouteRuleSetParam) (*view.PolicyRouteRuleSetInventoryView, error) {
	var resp view.UpdatePolicyRouteRuleSetEventView
	if err := cli.Put("v1/policy-routes/ruleSets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePolicyRouteRuleSet deletes PolicyRouteRuleSet
func (cli *ZSClient) DeletePolicyRouteRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/ruleSets/{uuid}", uuid, string(deleteMode))
}
// QueryPolicyRouteRuleSet queries PolicyRouteRuleSet list
func (cli *ZSClient) QueryPolicyRouteRuleSet(params *param.QueryParam) ([]view.PolicyRouteRuleSetInventoryView, error) {
	var resp []view.PolicyRouteRuleSetInventoryView
	return resp, cli.List("v1/policy-routes/rulesets", params, &resp)
}
