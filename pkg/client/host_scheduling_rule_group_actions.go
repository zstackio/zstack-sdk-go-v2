// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostSchedulingRuleGroup updates HostSchedulingRuleGroup
func (cli *ZSClient) UpdateHostSchedulingRuleGroup(uuid string, params param.UpdateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	resp := view.HostSchedulingRuleGroupInventoryView{}
	if err := cli.Put("v1/hostSchedulingRuleGroup", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostSchedulingRuleGroup queries HostSchedulingRuleGroup list
func (cli *ZSClient) QueryHostSchedulingRuleGroup(params *param.QueryParam) ([]view.HostSchedulingRuleGroupInventoryView, error) {
	var resp []view.HostSchedulingRuleGroupInventoryView
	return resp, cli.List("v1/query/host/schedulingRule/group", params, &resp)
}

func (cli *ZSClient) GetHostSchedulingRuleGroup(uuid string) (*view.HostSchedulingRuleGroupInventoryView, error) {
	var resp view.HostSchedulingRuleGroupInventoryView
	if err := cli.Get("v1/query/host/schedulingRule/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostSchedulingRuleGroup Pagination
func (cli *ZSClient) PageHostSchedulingRuleGroup(params *param.QueryParam) ([]view.HostSchedulingRuleGroupInventoryView, int, error) {
	var hostSchedulingRuleGroups []view.HostSchedulingRuleGroupInventoryView
	total, err := cli.Page("v1/query/host/schedulingRule/group", params, &hostSchedulingRuleGroups)
	return hostSchedulingRuleGroups, total, err
}
// CreateHostSchedulingRuleGroup creates HostSchedulingRuleGroup
func (cli *ZSClient) CreateHostSchedulingRuleGroup(params param.CreateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	resp := view.HostSchedulingRuleGroupInventoryView{}
	if err := cli.Post("v1/hostSchedulingRuleGroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteHostSchedulingRuleGroup deletes HostSchedulingRuleGroup
func (cli *ZSClient) DeleteHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hostSchedulingRuleGroup", uuid, string(deleteMode))
}
