// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryRole 查询Role列表
func (cli *ZSClient) QueryRole(params param.QueryParam) ([]view.QueryRoleView, error) {
	var resp []view.QueryRoleView
	return resp, cli.List("v1/identities/roles", &params, &resp)
}

