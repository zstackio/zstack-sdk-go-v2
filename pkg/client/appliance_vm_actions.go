// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryApplianceVm 查询ApplianceVm列表
func (cli *ZSClient) QueryApplianceVm(params param.QueryParam) ([]view.QueryApplianceVmView, error) {
	var resp []view.QueryApplianceVmView
	return resp, cli.List("v1/vm-instances/appliances", &params, &resp)
}

