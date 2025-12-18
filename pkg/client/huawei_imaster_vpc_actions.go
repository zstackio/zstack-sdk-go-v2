// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHuaweiIMasterVpc 查询HuaweiIMasterVpc列表
func (cli *ZSClient) QueryHuaweiIMasterVpc(params param.QueryParam) ([]view.QueryHuaweiIMasterVpcView, error) {
	var resp []view.QueryHuaweiIMasterVpcView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vpcs", &params, &resp)
}

