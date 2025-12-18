// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHostOsCategory 查询HostOsCategory列表
func (cli *ZSClient) QueryHostOsCategory(params param.QueryParam) ([]view.QueryHostOsCategoryView, error) {
	var resp []view.QueryHostOsCategoryView
	return resp, cli.List("v1/hosts/os/category", &params, &resp)
}

