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
	return cli.DeleteWithSpec("v1/vmSchedulingRuleGroup", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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
// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	var resp view.UpdateVmSchedulingRuleGroupEventView
	err := cli.PutWithSpec("v1/vmSchedulingRuleGroup", fmt.Sprintf(\"%s/update\", uuid), params, &resp)
	if err != nil {
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
