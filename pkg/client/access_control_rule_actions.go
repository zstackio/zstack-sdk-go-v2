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
	resp := view.AccessControlRuleInventoryView{}
	if err := cli.Post("v1/login-control/access-control/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAccessControlRule updates AccessControlRule
func (cli *ZSClient) UpdateAccessControlRule(uuid string, params param.UpdateAccessControlRuleParam) (*view.AccessControlRuleInventoryView, error) {
	resp := view.AccessControlRuleInventoryView{}
	if err := cli.Put("v1/login-control/access-control/rules", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccessControlRule deletes AccessControlRule
func (cli *ZSClient) DeleteAccessControlRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/login-control/access-control/rules", uuid, string(deleteMode))
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

// PageAccessControlRule Pagination
func (cli *ZSClient) PageAccessControlRule(params *param.QueryParam) ([]view.AccessControlRuleInventoryView, int, error) {
	var accessControlRules []view.AccessControlRuleInventoryView
	total, err := cli.Page("v1/login-control/access-control/rules", params, &accessControlRules)
	return accessControlRules, total, err
}
