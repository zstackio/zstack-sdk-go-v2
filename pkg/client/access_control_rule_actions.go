// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAccessControlRule adds AccessControlRule
func (cli *ZSClient) AddAccessControlRule(params param.AddAccessControlRuleParam) (*view.AccessControlRuleInventoryView, error) {
	var resp view.AddAccessControlRuleEventView
	if err := cli.Post("v1/login-control/access-control/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAccessControlRule updates AccessControlRule
func (cli *ZSClient) UpdateAccessControlRule(uuid string, params param.UpdateAccessControlRuleParam) (*view.AccessControlRuleInventoryView, error) {
	var resp view.UpdateAccessControlRuleEventView
	if err := cli.Put("v1/login-control/access-control/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAccessControlRule deletes AccessControlRule
func (cli *ZSClient) DeleteAccessControlRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/login-control/access-control/rules/{uuid}", uuid, string(deleteMode))
}
// QueryAccessControlRule queries AccessControlRule list
func (cli *ZSClient) QueryAccessControlRule(params *param.QueryParam) ([]view.AccessControlRuleInventoryView, error) {
	var resp []view.AccessControlRuleInventoryView
	return resp, cli.List("v1/login-control/access-control/rules", params, &resp)
}
