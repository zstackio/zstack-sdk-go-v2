// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveVmSchedulingRule removes VmSchedulingRule
func (cli *ZSClient) RemoveVmSchedulingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRule/{uuid}", uuid, string(deleteMode))
}
// CreateVmSchedulingRule creates VmSchedulingRule
func (cli *ZSClient) CreateVmSchedulingRule(params param.CreateVmSchedulingRuleParam) (*view.AffinityGroupInventoryView, error) {
	var resp view.CreateAffinityGroupEventView
	if err := cli.Post("v1/vmsSchedulingRule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ValidateVmSchedulingRule operates on VmSchedulingRule
func (cli *ZSClient) ValidateVmSchedulingRule(uuid string, params param.ValidateVmSchedulingRuleParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.Put("v1/validate/vmSchedulingRule", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmSchedulingRule updates VmSchedulingRule
func (cli *ZSClient) UpdateVmSchedulingRule(uuid string, params param.UpdateVmSchedulingRuleParam) (*view.VmSchedulingRuleInventoryView, error) {
	var resp view.UpdateVmSchedulingRuleEventView
	if err := cli.Put("v1/vmSchedulingRule/{uuid}/update", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVmSchedulingRule queries VmSchedulingRule list
func (cli *ZSClient) QueryVmSchedulingRule(params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, error) {
	var resp []view.VmSchedulingRuleInventoryView
	return resp, cli.List("v1/query/vm/schedulingRule", params, &resp)
}
