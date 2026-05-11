// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RemoveVmSchedulingRule removes VmSchedulingRule
func (cli *ZSClient) RemoveVmSchedulingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vmSchedulingRule", uuid, string(deleteMode))
}
// CreateVmSchedulingRule creates VmSchedulingRule
func (cli *ZSClient) CreateVmSchedulingRule(ctx context.Context, params param.CreateVmSchedulingRuleParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/vmsSchedulingRule", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateVmSchedulingRule operates on VmSchedulingRule
func (cli *ZSClient) ValidateVmSchedulingRule(ctx context.Context, params param.ValidateVmSchedulingRuleParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/validate/vmSchedulingRule", "", "", map[string]interface{}{
		"validateVmSchedulingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmSchedulingRule updates VmSchedulingRule
func (cli *ZSClient) UpdateVmSchedulingRule(ctx context.Context, uuid string, params param.UpdateVmSchedulingRuleParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vmSchedulingRule", uuid, "", map[string]interface{}{
		"updateVmSchedulingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmSchedulingRule queries VmSchedulingRule list
func (cli *ZSClient) QueryVmSchedulingRule(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, error) {
	var resp []view.VmSchedulingRuleInventoryView
	return resp, cli.List(ctx, "v1/query/vm/schedulingRule", params, &resp)
}

func (cli *ZSClient) GetVmSchedulingRule(ctx context.Context, uuid string) (*view.VmSchedulingRuleInventoryView, error) {
	var resp view.VmSchedulingRuleInventoryView
	if err := cli.Get(ctx, "v1/query/vm/schedulingRule", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmSchedulingRule Pagination
func (cli *ZSClient) PageVmSchedulingRule(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, int, error) {
	var vmSchedulingRules []view.VmSchedulingRuleInventoryView
	total, err := cli.Page(ctx, "v1/query/vm/schedulingRule", params, &vmSchedulingRules)
	return vmSchedulingRules, total, err
}
