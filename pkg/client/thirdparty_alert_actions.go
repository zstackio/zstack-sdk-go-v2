// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryThirdpartyAlert 查询ThirdpartyAlert列表
func (cli *ZSClient) QueryThirdpartyAlert(params param.QueryParam) ([]view.QueryThirdpartyAlertView, error) {
	var resp []view.QueryThirdpartyAlertView
	return resp, cli.List("v1/zwatch/third-party/alerts", &params, &resp)
}

