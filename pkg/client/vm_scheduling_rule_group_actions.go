// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVmSchedulingRuleGroup deletes VmSchedulingRuleGroup
func (cli *ZSClient) DeleteVmSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vmSchedulingRuleGroup", uuid, string(deleteMode))
}
// QueryVmSchedulingRuleGroup queries VmSchedulingRuleGroup list
func (cli *ZSClient) QueryVmSchedulingRuleGroup(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, error) {
	var resp []view.VmSchedulingRuleGroupInventoryView
	return resp, cli.List(ctx, "v1/query/vm/schedulingRule/group", params, &resp)
}

func (cli *ZSClient) GetVmSchedulingRuleGroup(ctx context.Context, uuid string) (*view.VmSchedulingRuleGroupInventoryView, error) {
	var resp view.VmSchedulingRuleGroupInventoryView
	if err := cli.Get(ctx, "v1/query/vm/schedulingRule/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmSchedulingRuleGroup Pagination
func (cli *ZSClient) PageVmSchedulingRuleGroup(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, int, error) {
	var vmSchedulingRuleGroups []view.VmSchedulingRuleGroupInventoryView
	total, err := cli.Page(ctx, "v1/query/vm/schedulingRule/group", params, &vmSchedulingRuleGroups)
	return vmSchedulingRuleGroups, total, err
}
// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(ctx context.Context, uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	resp := view.VmSchedulingRuleGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vmSchedulingRuleGroup", uuid, "", map[string]interface{}{
		"updateVmSchedulingRuleGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVmSchedulingRuleGroup creates VmSchedulingRuleGroup
func (cli *ZSClient) CreateVmSchedulingRuleGroup(ctx context.Context, params param.CreateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	resp := view.VmSchedulingRuleGroupInventoryView{}
	if err := cli.Post(ctx, "v1/vmSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
