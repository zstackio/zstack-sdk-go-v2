// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestToolsState 查询GuestToolsState列表
func (cli *ZSClient) QueryGuestToolsState(params param.QueryParam) ([]view.QueryGuestToolsStateView, error) {
	var resp []view.QueryGuestToolsStateView
	return resp, cli.List("v1/guesttools", &params, &resp)
}

