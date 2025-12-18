// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryL2Network 查询L2Network列表
func (cli *ZSClient) QueryL2Network(params param.QueryParam) ([]view.QueryL2NetworkView, error) {
	var resp []view.QueryL2NetworkView
	return resp, cli.List("v1/l2-networks", &params, &resp)
}

