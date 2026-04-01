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
	resp := view.PolicyRouteRuleSetInventoryView{}
	if err := cli.Post("v1/policy-routes/rulesets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePolicyRouteRuleSet updates PolicyRouteRuleSet
func (cli *ZSClient) UpdatePolicyRouteRuleSet(uuid string, params param.UpdatePolicyRouteRuleSetParam) (*view.PolicyRouteRuleSetInventoryView, error) {
	resp := view.PolicyRouteRuleSetInventoryView{}
	if err := cli.PutWithRespKey("v1/policy-routes/ruleSets", uuid, "", map[string]interface{}{
		"updatePolicyRouteRuleSet": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteRuleSet deletes PolicyRouteRuleSet
func (cli *ZSClient) DeletePolicyRouteRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/ruleSets", uuid, string(deleteMode))
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

// PagePolicyRouteRuleSet Pagination
func (cli *ZSClient) PagePolicyRouteRuleSet(params *param.QueryParam) ([]view.PolicyRouteRuleSetInventoryView, int, error) {
	var policyRouteRuleSets []view.PolicyRouteRuleSetInventoryView
	total, err := cli.Page("v1/policy-routes/rulesets", params, &policyRouteRuleSets)
	return policyRouteRuleSets, total, err
}
