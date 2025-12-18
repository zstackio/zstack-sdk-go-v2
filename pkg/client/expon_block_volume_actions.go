// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryExponBlockVolume 查询ExponBlockVolume列表
func (cli *ZSClient) QueryExponBlockVolume(params param.QueryParam) ([]view.QueryExponBlockVolumeView, error) {
	var resp []view.QueryExponBlockVolumeView
	return resp, cli.List("v1/expon/block-volumes", &params, &resp)
}

