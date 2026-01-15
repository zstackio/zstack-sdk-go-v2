// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVmSchedulingRuleGroup deletes VmSchedulingRuleGroup
func (cli *ZSClient) DeleteVmSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRuleGroup", uuid, string(deleteMode))
}
// QueryVmSchedulingRuleGroup queries VmSchedulingRuleGroup list
func (cli *ZSClient) QueryVmSchedulingRuleGroup(params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, error) {
	var resp []view.VmSchedulingRuleGroupInventoryView
	return resp, cli.List("v1/query/vm/schedulingRule/group", params, &resp)
}

func (cli *ZSClient) GetVmSchedulingRuleGroup(uuid string) (*view.VmSchedulingRuleGroupInventoryView, error) {
	var resp view.VmSchedulingRuleGroupInventoryView
	if err := cli.Get("v1/query/vm/schedulingRule/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmSchedulingRuleGroup Pagination
func (cli *ZSClient) PageVmSchedulingRuleGroup(params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, int, error) {
	var vmSchedulingRuleGroups []view.VmSchedulingRuleGroupInventoryView
	total, err := cli.Page("v1/query/vm/schedulingRule/group", params, &vmSchedulingRuleGroups)
	return vmSchedulingRuleGroups, total, err
}
// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	resp := view.VmSchedulingRuleGroupInventoryView{}
	if err := cli.Put("v1/vmSchedulingRuleGroup", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVmSchedulingRuleGroup creates VmSchedulingRuleGroup
func (cli *ZSClient) CreateVmSchedulingRuleGroup(params param.CreateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	resp := view.VmSchedulingRuleGroupInventoryView{}
	if err := cli.Post("v1/vmSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
