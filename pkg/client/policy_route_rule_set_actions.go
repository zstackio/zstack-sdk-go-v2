// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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
	err := cli.PutWithSpec("v1/policy-routes/ruleSets", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePolicyRouteRuleSet deletes PolicyRouteRuleSet
func (cli *ZSClient) DeletePolicyRouteRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/policy-routes/ruleSets", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryPolicyRouteRuleSet queries PolicyRouteRuleSet list
func (cli *ZSClient) QueryPolicyRouteRuleSet(params *param.QueryParam) ([]view.PolicyRouteRuleSetInventoryView, error) {
	var resp []view.PolicyRouteRuleSetInventoryView
	return resp, cli.List("v1/policy-routes/rulesets", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRuleSet(uuid string) (*view.PolicyRouteRuleSetInventoryView, error) {
	var resp view.PolicyRouteRuleSetInventoryView
	if err := cli.Get("v1/policy-routes/rulesets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
