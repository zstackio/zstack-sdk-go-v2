// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePolicy deletes Policy
func (cli *ZSClient) DeletePolicy(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/policies/{uuid}", uuid, string(deleteMode))
}
// QueryPolicy queries Policy list
func (cli *ZSClient) QueryPolicy(params *param.QueryParam) ([]view.PolicyInventoryView, error) {
	var resp []view.PolicyInventoryView
	return resp, cli.List("v1/accounts/policies", params, &resp)
}
// CreatePolicy creates Policy
func (cli *ZSClient) CreatePolicy(params param.CreatePolicyParam) (*view.PolicyInventoryView, error) {
	var resp view.CreatePolicyEventView
	if err := cli.Post("v1/accounts/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
