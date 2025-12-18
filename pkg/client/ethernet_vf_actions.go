// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEthernetVF 查询EthernetVF列表
func (cli *ZSClient) QueryEthernetVF(params param.QueryParam) ([]view.QueryEthernetVFView, error) {
	var resp []view.QueryEthernetVFView
	return resp, cli.List("v1/pci-device/ethernet-vfs", &params, &resp)
}

