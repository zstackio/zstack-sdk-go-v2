// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHostPhysicalMemory 查询HostPhysicalMemory列表
func (cli *ZSClient) QueryHostPhysicalMemory(params param.QueryParam) ([]view.QueryHostPhysicalMemoryView, error) {
	var resp []view.QueryHostPhysicalMemoryView
	return resp, cli.List("v1/hosts/physicalmemory", &params, &resp)
}

