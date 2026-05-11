// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostSchedulingRuleGroup updates HostSchedulingRuleGroup
func (cli *ZSClient) UpdateHostSchedulingRuleGroup(ctx context.Context, uuid string, params param.UpdateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	resp := view.HostSchedulingRuleGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hostSchedulingRuleGroup", uuid, "", map[string]interface{}{
		"updateHostSchedulingRuleGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostSchedulingRuleGroup queries HostSchedulingRuleGroup list
func (cli *ZSClient) QueryHostSchedulingRuleGroup(ctx context.Context, params *param.QueryParam) ([]view.HostSchedulingRuleGroupInventoryView, error) {
	var resp []view.HostSchedulingRuleGroupInventoryView
	return resp, cli.List(ctx, "v1/query/host/schedulingRule/group", params, &resp)
}

func (cli *ZSClient) GetHostSchedulingRuleGroup(ctx context.Context, uuid string) (*view.HostSchedulingRuleGroupInventoryView, error) {
	var resp view.HostSchedulingRuleGroupInventoryView
	if err := cli.Get(ctx, "v1/query/host/schedulingRule/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostSchedulingRuleGroup Pagination
func (cli *ZSClient) PageHostSchedulingRuleGroup(ctx context.Context, params *param.QueryParam) ([]view.HostSchedulingRuleGroupInventoryView, int, error) {
	var hostSchedulingRuleGroups []view.HostSchedulingRuleGroupInventoryView
	total, err := cli.Page(ctx, "v1/query/host/schedulingRule/group", params, &hostSchedulingRuleGroups)
	return hostSchedulingRuleGroups, total, err
}
// CreateHostSchedulingRuleGroup creates HostSchedulingRuleGroup
func (cli *ZSClient) CreateHostSchedulingRuleGroup(ctx context.Context, params param.CreateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	resp := view.HostSchedulingRuleGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/hostSchedulingRuleGroup", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteHostSchedulingRuleGroup deletes HostSchedulingRuleGroup
func (cli *ZSClient) DeleteHostSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hostSchedulingRuleGroup", uuid, string(deleteMode))
}
