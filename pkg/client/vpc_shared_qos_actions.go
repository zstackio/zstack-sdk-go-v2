// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcSharedQos 查询VpcSharedQos列表
func (cli *ZSClient) QueryVpcSharedQos(params param.QueryParam) ([]view.QueryVpcSharedQosView, error) {
	var resp []view.QueryVpcSharedQosView
	return resp, cli.List("v1/vips/sharedqos", &params, &resp)
}

