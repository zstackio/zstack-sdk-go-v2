// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryGuestToolsState queries GuestToolsState list
func (cli *ZSClient) QueryGuestToolsState(params *param.QueryParam) ([]view.GuestToolsStateInventoryView, error) {
	var resp []view.GuestToolsStateInventoryView
	return resp, cli.List("v1/guesttools", params, &resp)
}
