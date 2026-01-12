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
	var resp view.UpdateHostSchedulingRuleGroupEventView
	err := cli.PutWithSpec("v1/hostSchedulingRuleGroup", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
	return cli.DeleteWithSpec("v1/hostSchedulingRuleGroup", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
