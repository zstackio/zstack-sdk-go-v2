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

// PagePolicy Pagination
func (cli *ZSClient) PagePolicy(params *param.QueryParam) ([]view.PolicyInventoryView, int, error) {
	var policies []view.PolicyInventoryView
	total, err := cli.Page("v1/accounts/policies", params, &policies)
	return policies, total, err
}
// CreatePolicy creates Policy
func (cli *ZSClient) CreatePolicy(params param.CreatePolicyParam) (*view.PolicyInventoryView, error) {
	resp := view.PolicyInventoryView{}
	if err := cli.Post("v1/accounts/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
