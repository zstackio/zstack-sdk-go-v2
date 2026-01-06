// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePortForwardingRule deletes PortForwardingRule
func (cli *ZSClient) DeletePortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding/{uuid}", uuid, string(deleteMode))
}
// QueryPortForwardingRule queries PortForwardingRule list
func (cli *ZSClient) QueryPortForwardingRule(params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, error) {
	var resp []view.PortForwardingRuleInventoryView
	return resp, cli.List("v1/port-forwarding", params, &resp)
}
// UpdatePortForwardingRule updates PortForwardingRule
func (cli *ZSClient) UpdatePortForwardingRule(uuid string, params param.UpdatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.UpdatePortForwardingRuleEventView
	if err := cli.Put("v1/port-forwarding/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DetachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) DetachPortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding/{uuid}/vm-instances/nics", uuid, string(deleteMode))
}
// AttachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) AttachPortForwardingRule(params param.AttachPortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.AttachPortForwardingRuleEventView
	if err := cli.Post("v1/port-forwarding/{ruleUuid}/vm-instances/nics/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreatePortForwardingRule creates PortForwardingRule
func (cli *ZSClient) CreatePortForwardingRule(params param.CreatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.CreatePortForwardingRuleEventView
	if err := cli.Post("v1/port-forwarding", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
