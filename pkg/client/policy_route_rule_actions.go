// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPolicyRouteRule queries PolicyRouteRule list
func (cli *ZSClient) QueryPolicyRouteRule(params *param.QueryParam) ([]view.PolicyRouteRuleInventoryView, error) {
	var resp []view.PolicyRouteRuleInventoryView
	return resp, cli.List("v1/policy-routes/rules", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRule(uuid string) (*view.PolicyRouteRuleInventoryView, error) {
	var resp view.PolicyRouteRuleInventoryView
	if err := cli.Get("v1/policy-routes/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteRule Pagination
func (cli *ZSClient) PagePolicyRouteRule(params *param.QueryParam) ([]view.PolicyRouteRuleInventoryView, int, error) {
	var policyRouteRules []view.PolicyRouteRuleInventoryView
	total, err := cli.Page("v1/policy-routes/rules", params, &policyRouteRules)
	return policyRouteRules, total, err
}
// CreatePolicyRouteRule creates PolicyRouteRule
func (cli *ZSClient) CreatePolicyRouteRule(params param.CreatePolicyRouteRuleParam) (*view.PolicyRouteRuleInventoryView, error) {
	resp := view.PolicyRouteRuleInventoryView{}
	if err := cli.Post("v1/policy-routes/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteRule deletes PolicyRouteRule
func (cli *ZSClient) DeletePolicyRouteRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/rules", uuid, string(deleteMode))
}
