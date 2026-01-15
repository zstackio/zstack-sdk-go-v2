// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePortForwardingRule deletes PortForwardingRule
func (cli *ZSClient) DeletePortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding", uuid, string(deleteMode))
}
// QueryPortForwardingRule queries PortForwardingRule list
func (cli *ZSClient) QueryPortForwardingRule(params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, error) {
	var resp []view.PortForwardingRuleInventoryView
	return resp, cli.List("v1/port-forwarding", params, &resp)
}

// PagePortForwardingRule Pagination
func (cli *ZSClient) PagePortForwardingRule(params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, int, error) {
	var portForwardingRules []view.PortForwardingRuleInventoryView
	total, err := cli.Page("v1/port-forwarding", params, &portForwardingRules)
	return portForwardingRules, total, err
}
// UpdatePortForwardingRule updates PortForwardingRule
func (cli *ZSClient) UpdatePortForwardingRule(uuid string, params param.UpdatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.Put("v1/port-forwarding", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DetachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) DetachPortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding", uuid, string(deleteMode))
}
// AttachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) AttachPortForwardingRule(ruleUuid string, vmNicUuid string, params param.AttachPortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	err := cli.Post(fmt.Sprintf("v1/port-forwarding/%s/vm-instances/nics/%s", ruleUuid, vmNicUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreatePortForwardingRule creates PortForwardingRule
func (cli *ZSClient) CreatePortForwardingRule(params param.CreatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.Post("v1/port-forwarding", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
