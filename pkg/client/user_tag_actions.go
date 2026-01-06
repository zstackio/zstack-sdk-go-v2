// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
