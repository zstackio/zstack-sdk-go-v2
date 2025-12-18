// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventFromResourceStack 查询EventFromResourceStack列表
func (cli *ZSClient) QueryEventFromResourceStack(params param.QueryParam) ([]view.QueryEventFromResourceStackView, error) {
	var resp []view.QueryEventFromResourceStackView
	return resp, cli.List("v1/cloudformation/event", &params, &resp)
}

