// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostSchedulingRuleGroup updates HostSchedulingRuleGroup
func (cli *ZSClient) UpdateHostSchedulingRuleGroup(uuid string, params param.UpdateHostSchedulingRuleGroupParam) (*view.HostSchedulingRuleGroupInventoryView, error) {
	resp := view.HostSchedulingRuleGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/hostSchedulingRuleGroup", uuid, "", map[string]interface{}{
		"updateHostSchedulingRuleGroup": params.Params,
	}, &resp); err != nil {
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

// AddHostToHostSchedulingRuleGroup adds HostToHostSchedulingRuleGroup
func (cli *ZSClient) AddHostToHostSchedulingRuleGroup(hostGroupUuid string, hostUuid string, params param.AddHostToHostSchedulingRuleGroupParam) (*view.AddHostToHostSchedulingRuleGroupEventView, error) {
	resp := view.AddHostToHostSchedulingRuleGroupEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hostSchedulingRuleGroup/%s/host/%s", hostGroupUuid, hostUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHostFromHostSchedulingRuleGroup requires additional arguments.
//
// Deprecated: use DetachHostFromHostSchedulingRuleGroupWithParams to supply the required fields.
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroup(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DetachHostFromHostSchedulingRuleGroup requires additional arguments; use DetachHostFromHostSchedulingRuleGroupWithParams")
}

// DetachHostFromHostSchedulingRuleGroupWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroupWithParams(hostGroupUuid string, deleteMode param.DeleteMode, params param.DetachHostFromHostSchedulingRuleGroupParamDetail) error {
	query := url.Values{
		"hostUuid": {params.HostUuid},
	}
	return cli.deleteActionWithQuery("v1/hostSchedulingRuleGroup", hostGroupUuid, "host", deleteMode, query, "hostUuid")
}
