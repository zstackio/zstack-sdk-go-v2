// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAccount queries Account list
func (cli *ZSClient) QueryAccount(params *param.QueryParam) ([]view.AccountInventoryView, error) {
	var resp []view.AccountInventoryView
	return resp, cli.List("v1/accounts", params, &resp)
}
