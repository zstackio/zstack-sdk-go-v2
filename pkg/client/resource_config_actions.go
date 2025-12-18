// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryResourceConfig 查询ResourceConfig列表
func (cli *ZSClient) QueryResourceConfig(params param.QueryParam) ([]view.QueryResourceConfigView, error) {
	var resp []view.QueryResourceConfigView
	return resp, cli.List("v1/resource-configurations", &params, &resp)
}

