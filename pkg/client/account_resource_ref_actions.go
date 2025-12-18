// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccountResourceRef 查询AccountResourceRef列表
func (cli *ZSClient) QueryAccountResourceRef(params param.QueryParam) ([]view.QueryAccountResourceRefView, error) {
	var resp []view.QueryAccountResourceRefView
	return resp, cli.List("v1/accounts/resources/refs", &params, &resp)
}

