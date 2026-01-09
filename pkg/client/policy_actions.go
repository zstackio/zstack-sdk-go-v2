// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePolicy deletes Policy
func (cli *ZSClient) DeletePolicy(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/policies", uuid, string(deleteMode))
}
// QueryPolicy queries Policy list
func (cli *ZSClient) QueryPolicy(params *param.QueryParam) ([]view.PolicyInventoryView, error) {
	var resp []view.PolicyInventoryView
	return resp, cli.List("v1/accounts/policies", params, &resp)
}

func (cli *ZSClient) GetPolicy(uuid string) (*view.PolicyInventoryView, error) {
	var resp view.PolicyInventoryView
	if err := cli.Get("v1/accounts/policies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreatePolicy creates Policy
func (cli *ZSClient) CreatePolicy(params param.CreatePolicyParam) (*view.PolicyInventoryView, error) {
	var resp view.CreatePolicyEventView
	if err := cli.Post("v1/accounts/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
