// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryContainerImage 查询ContainerImage列表
func (cli *ZSClient) QueryContainerImage(params param.QueryParam) ([]view.QueryContainerImageView, error) {
	var resp []view.QueryContainerImageView
	return resp, cli.List("v1/container/images", &params, &resp)
}

