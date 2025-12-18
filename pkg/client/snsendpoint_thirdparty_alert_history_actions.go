// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSEndpointThirdpartyAlertHistory 查询SNSEndpointThirdpartyAlertHistory列表
func (cli *ZSClient) QuerySNSEndpointThirdpartyAlertHistory(params param.QueryParam) ([]view.QuerySNSEndpointThirdpartyAlertHistoryView, error) {
	var resp []view.QuerySNSEndpointThirdpartyAlertHistoryView
	return resp, cli.List("v1/zwatch/third-party/alert-publish-histories", &params, &resp)
}

