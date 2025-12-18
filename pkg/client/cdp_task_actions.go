// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCdpTask 查询CdpTask列表
func (cli *ZSClient) QueryCdpTask(params param.QueryParam) ([]view.QueryCdpTaskView, error) {
	var resp []view.QueryCdpTaskView
	return resp, cli.List("v1/cdp-task", &params, &resp)
}

