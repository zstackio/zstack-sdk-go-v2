// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryUserGroup queries UserGroup list
func (cli *ZSClient) QueryUserGroup(params *param.QueryParam) ([]view.UserGroupInventoryView, error) {
	var resp []view.UserGroupInventoryView
	return resp, cli.List("v1/accounts/groups", params, &resp)
}
