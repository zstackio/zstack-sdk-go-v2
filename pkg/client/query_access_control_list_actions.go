// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAccessControlList queries AccessControlList list
func (cli *ZSClient) QueryAccessControlList(params *param.QueryParam) ([]view.AccessControlListInventoryView, error) {
	var resp []view.AccessControlListInventoryView
	return resp, cli.List("v1/access-control-lists", params, &resp)
}
