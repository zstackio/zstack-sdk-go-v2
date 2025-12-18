// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNativeHost 查询NativeHost列表
func (cli *ZSClient) QueryNativeHost(params param.QueryParam) ([]view.QueryNativeHostView, error) {
	var resp []view.QueryNativeHostView
	return resp, cli.List("v1/container/native/host", &params, &resp)
}

