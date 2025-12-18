// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMttyDevice 查询MttyDevice列表
func (cli *ZSClient) QueryMttyDevice(params param.QueryParam) ([]view.QueryMttyDeviceView, error) {
	var resp []view.QueryMttyDeviceView
	return resp, cli.List("v1/mtty-devices", &params, &resp)
}

