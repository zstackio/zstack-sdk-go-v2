// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAccessControlList creates AccessControlList
func (cli *ZSClient) CreateAccessControlList(ctx context.Context, params param.CreateAccessControlListParam) (*view.AccessControlListInventoryView, error) {
	resp := view.AccessControlListInventoryView{}
	if err := cli.Post(ctx, "v1/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAccessControlList updates AccessControlList
func (cli *ZSClient) UpdateAccessControlList(ctx context.Context, uuid string, params param.UpdateAccessControlListParam) (*view.AccessControlListInventoryView, error) {
	resp := view.AccessControlListInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/access-control-lists", uuid, "", map[string]interface{}{
		"updateAccessControlList": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAccessControlList queries AccessControlList list
func (cli *ZSClient) QueryAccessControlList(ctx context.Context, params *param.QueryParam) ([]view.AccessControlListInventoryView, error) {
	var resp []view.AccessControlListInventoryView
	return resp, cli.List(ctx, "v1/access-control-lists", params, &resp)
}

func (cli *ZSClient) GetAccessControlList(ctx context.Context, uuid string) (*view.AccessControlListInventoryView, error) {
	var resp view.AccessControlListInventoryView
	if err := cli.Get(ctx, "v1/access-control-lists", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccessControlList Pagination
func (cli *ZSClient) PageAccessControlList(ctx context.Context, params *param.QueryParam) ([]view.AccessControlListInventoryView, int, error) {
	var accessControlLists []view.AccessControlListInventoryView
	total, err := cli.Page(ctx, "v1/access-control-lists", params, &accessControlLists)
	return accessControlLists, total, err
}
// DeleteAccessControlList deletes AccessControlList
func (cli *ZSClient) DeleteAccessControlList(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/access-control-lists", uuid, string(deleteMode))
}
