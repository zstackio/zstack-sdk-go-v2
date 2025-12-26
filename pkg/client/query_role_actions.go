// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryRole queries Role list
func (cli *ZSClient) QueryRole(params *param.QueryParam) ([]view.RoleInventoryView, error) {
	var resp []view.RoleInventoryView
	return resp, cli.List("v1/identities/roles", params, &resp)
}
