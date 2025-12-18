// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryImageGroupRef 查询ImageGroupRef列表
func (cli *ZSClient) QueryImageGroupRef(params param.QueryParam) ([]view.QueryImageGroupRefView, error) {
	var resp []view.QueryImageGroupRefView
	return resp, cli.List("v1/imagegrouprefs", &params, &resp)
}

