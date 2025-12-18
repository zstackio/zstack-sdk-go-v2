// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAlertDataAck 查询AlertDataAck列表
func (cli *ZSClient) QueryAlertDataAck(params param.QueryParam) ([]view.QueryAlertDataAckView, error) {
	var resp []view.QueryAlertDataAckView
	return resp, cli.List("v1/zwatch/alert-histories/acknowledgments", &params, &resp)
}

