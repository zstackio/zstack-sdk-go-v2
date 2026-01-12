// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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
	err := cli.PutWithSpec("v1/login-control/access-control/rules", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAccessControlRule deletes AccessControlRule
func (cli *ZSClient) DeleteAccessControlRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/login-control/access-control/rules", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryAccessControlRule queries AccessControlRule list
func (cli *ZSClient) QueryAccessControlRule(params *param.QueryParam) ([]view.AccessControlRuleInventoryView, error) {
	var resp []view.AccessControlRuleInventoryView
	return resp, cli.List("v1/login-control/access-control/rules", params, &resp)
}

func (cli *ZSClient) GetAccessControlRule(uuid string) (*view.AccessControlRuleInventoryView, error) {
	var resp view.AccessControlRuleInventoryView
	if err := cli.Get("v1/login-control/access-control/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
