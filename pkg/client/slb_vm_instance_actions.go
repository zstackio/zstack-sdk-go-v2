// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySlbVmInstance 查询SlbVmInstance列表
func (cli *ZSClient) QuerySlbVmInstance(params param.QueryParam) ([]view.QuerySlbVmInstanceView, error) {
	var resp []view.QuerySlbVmInstanceView
	return resp, cli.List("v1/load-balancers/slb/instances", &params, &resp)
}

