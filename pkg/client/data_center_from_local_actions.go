// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDataCenterFromLocal 查询DataCenterFromLocal列表
func (cli *ZSClient) QueryDataCenterFromLocal(params param.QueryParam) ([]view.QueryDataCenterFromLocalView, error) {
	var resp []view.QueryDataCenterFromLocalView
	return resp, cli.List("v1/hybrid/data-center", &params, &resp)
}

