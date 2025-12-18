// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestVmScript 查询GuestVmScript列表
func (cli *ZSClient) QueryGuestVmScript(params param.QueryParam) ([]view.QueryGuestVmScriptView, error) {
	var resp []view.QueryGuestVmScriptView
	return resp, cli.List("v1/scripts", &params, &resp)
}

