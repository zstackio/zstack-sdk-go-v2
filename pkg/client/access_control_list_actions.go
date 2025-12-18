// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccessControlList 查询AccessControlList列表
func (cli *ZSClient) QueryAccessControlList(params param.QueryParam) ([]view.QueryAccessControlListView, error) {
	var resp []view.QueryAccessControlListView
	return resp, cli.List("v1/access-control-lists", &params, &resp)
}

