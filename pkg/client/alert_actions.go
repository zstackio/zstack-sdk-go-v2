// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAlert 查询Alert列表
func (cli *ZSClient) QueryAlert(params param.QueryParam) ([]view.QueryAlertView, error) {
	var resp []view.QueryAlertView
	return resp, cli.List("v1/monitoring/alerts", &params, &resp)
}

