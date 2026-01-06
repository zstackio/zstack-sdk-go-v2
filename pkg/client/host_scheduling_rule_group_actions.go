// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostSchedulingRuleGroup updates HostSchedulingRuleGroup
func (cli *ZSClient) UpdateHostSchedulingRuleGroup(uuid string, params param.UpdateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	var resp view.UpdateHostSchedulingRuleGroupEventView
	if err := cli.Put("v1/hostSchedulingRuleGroup/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryHostSchedulingRuleGroup queries HostSchedulingRuleGroup list
func (cli *ZSClient) QueryHostSchedulingRuleGroup(params *param.QueryParam) ([]view.HostSchedulingRuleGroupInventoryView, error) {
	var resp []view.HostSchedulingRuleGroupInventoryView
	return resp, cli.List("v1/query/host/schedulingRule/group", params, &resp)
}
// CreateHostSchedulingRuleGroup creates HostSchedulingRuleGroup
func (cli *ZSClient) CreateHostSchedulingRuleGroup(params param.CreateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	var resp view.CreateHostSchedulingRuleGroupEventView
	if err := cli.Post("v1/hostSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteHostSchedulingRuleGroup deletes HostSchedulingRuleGroup
func (cli *ZSClient) DeleteHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hostSchedulingRuleGroup/{uuid}", uuid, string(deleteMode))
}
