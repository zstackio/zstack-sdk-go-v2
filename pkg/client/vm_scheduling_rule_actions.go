// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveVmSchedulingRule removes VmSchedulingRule
func (cli *ZSClient) RemoveVmSchedulingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vmSchedulingRule", uuid, string(deleteMode))
}
// CreateVmSchedulingRule creates VmSchedulingRule
func (cli *ZSClient) CreateVmSchedulingRule(params param.CreateVmSchedulingRuleParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Post("v1/vmsSchedulingRule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.Put("v1/vmSchedulingRule", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmSchedulingRule queries VmSchedulingRule list
func (cli *ZSClient) QueryVmSchedulingRule(params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, error) {
	var resp []view.VmSchedulingRuleInventoryView
	return resp, cli.List("v1/query/vm/schedulingRule", params, &resp)
}

func (cli *ZSClient) GetVmSchedulingRule(uuid string) (*view.VmSchedulingRuleInventoryView, error) {
	var resp view.VmSchedulingRuleInventoryView
	if err := cli.Get("v1/query/vm/schedulingRule", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmSchedulingRule Pagination
func (cli *ZSClient) PageVmSchedulingRule(params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, int, error) {
	var vmSchedulingRules []view.VmSchedulingRuleInventoryView
	total, err := cli.Page("v1/query/vm/schedulingRule", params, &vmSchedulingRules)
	return vmSchedulingRules, total, err
}
