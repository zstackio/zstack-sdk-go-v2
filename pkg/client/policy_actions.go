// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePolicy deletes Policy
func (cli *ZSClient) DeletePolicy(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts/policies", uuid, string(deleteMode))
}
// QueryPolicy queries Policy list
func (cli *ZSClient) QueryPolicy(ctx context.Context, params *param.QueryParam) ([]view.PolicyInventoryView, error) {
	var resp []view.PolicyInventoryView
	return resp, cli.List(ctx, "v1/accounts/policies", params, &resp)
}

func (cli *ZSClient) GetPolicy(ctx context.Context, uuid string) (*view.PolicyInventoryView, error) {
	var resp view.PolicyInventoryView
	if err := cli.Get(ctx, "v1/accounts/policies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicy Pagination
func (cli *ZSClient) PagePolicy(ctx context.Context, params *param.QueryParam) ([]view.PolicyInventoryView, int, error) {
	var policies []view.PolicyInventoryView
	total, err := cli.Page(ctx, "v1/accounts/policies", params, &policies)
	return policies, total, err
}
// CreatePolicy creates Policy
func (cli *ZSClient) CreatePolicy(ctx context.Context, params param.CreatePolicyParam) (*view.PolicyInventoryView, error) {
	resp := view.PolicyInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/accounts/policies", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
