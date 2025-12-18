// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFcHbaDevice 查询FcHbaDevice列表
func (cli *ZSClient) QueryFcHbaDevice(params param.QueryParam) ([]view.QueryFcHbaDeviceView, error) {
	var resp []view.QueryFcHbaDeviceView
	return resp, cli.List("v1/storage-devices/hba", &params, &resp)
}

