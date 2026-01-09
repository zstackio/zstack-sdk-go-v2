// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateUserTag creates UserTag
func (cli *ZSClient) CreateUserTag(params param.CreateUserTagParam) (*view.UserTagInventoryView, error) {
	var resp view.CreateUserTagEventView
	if err := cli.Post("v1/user-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryUserTag queries UserTag list
func (cli *ZSClient) QueryUserTag(params *param.QueryParam) ([]view.UserTagInventoryView, error) {
	var resp []view.UserTagInventoryView
	return resp, cli.List("v1/user-tags", params, &resp)
}

func (cli *ZSClient) GetUserTag(uuid string) (*view.UserTagInventoryView, error) {
	var resp view.UserTagInventoryView
	if err := cli.Get("v1/user-tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
