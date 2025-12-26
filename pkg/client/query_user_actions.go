// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryUser queries User list
func (cli *ZSClient) QueryUser(params *param.QueryParam) ([]view.UserInventoryView, error) {
	var resp []view.UserInventoryView
	return resp, cli.List("v1/accounts/users", params, &resp)
}
