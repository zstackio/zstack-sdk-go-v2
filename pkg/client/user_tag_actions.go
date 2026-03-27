// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateUserTag creates UserTag
func (cli *ZSClient) CreateUserTag(ctx context.Context) (*view.UserTagInventoryView, error) {
	resp := view.UserTagInventoryView{}
	if err := cli.Post(ctx, "v1/user-tags", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryUserTag queries UserTag list
func (cli *ZSClient) QueryUserTag(ctx context.Context, params *param.QueryParam) ([]view.UserTagInventoryView, error) {
	var resp []view.UserTagInventoryView
	return resp, cli.List(ctx, "v1/user-tags", params, &resp)
}

func (cli *ZSClient) GetUserTag(ctx context.Context, uuid string) (*view.UserTagInventoryView, error) {
	var resp view.UserTagInventoryView
	if err := cli.Get(ctx, "v1/user-tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUserTag Pagination
func (cli *ZSClient) PageUserTag(ctx context.Context, params *param.QueryParam) ([]view.UserTagInventoryView, int, error) {
	var userTags []view.UserTagInventoryView
	total, err := cli.Page(ctx, "v1/user-tags", params, &userTags)
	return userTags, total, err
}
