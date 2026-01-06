// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVmSchedulingRuleGroup deletes VmSchedulingRuleGroup
func (cli *ZSClient) DeleteVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup/{uuid}", uuid, string(deleteMode))
}
// QueryVmSchedulingRuleGroup queries VmSchedulingRuleGroup list
func (cli *ZSClient) QueryVmSchedulingRuleGroup(params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, error) {
	var resp []view.VmSchedulingRuleGroupInventoryView
	return resp, cli.List("v1/query/vm/schedulingRule/group", params, &resp)
}
// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	var resp view.UpdateVmSchedulingRuleGroupEventView
	if err := cli.Put("v1/vmSchedulingRuleGroup/{uuid}/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateVmSchedulingRuleGroup creates VmSchedulingRuleGroup
func (cli *ZSClient) CreateVmSchedulingRuleGroup(params param.CreateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	var resp view.CreateVmSchedulingRuleGroupEventView
	if err := cli.Post("v1/vmSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
