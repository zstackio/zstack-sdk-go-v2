// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryGuestVmScript queries GuestVmScript list
func (cli *ZSClient) QueryGuestVmScript(params *param.QueryParam) ([]view.GuestVmScriptInventoryView, error) {
	var resp []view.GuestVmScriptInventoryView
	return resp, cli.List("v1/scripts", params, &resp)
}
