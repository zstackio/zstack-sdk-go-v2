// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryUserTag queries UserTag list
func (cli *ZSClient) QueryUserTag(params *param.QueryParam) ([]view.UserTagInventoryView, error) {
	var resp []view.UserTagInventoryView
	return resp, cli.List("v1/user-tags", params, &resp)
}
