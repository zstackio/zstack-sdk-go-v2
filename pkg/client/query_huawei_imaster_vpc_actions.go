// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHuaweiIMasterVpc queries HuaweiIMasterVpc list
func (cli *ZSClient) QueryHuaweiIMasterVpc(params param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, error) {
	var resp []view.HuaweiIMasterVpcInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vpcs", &params, &resp)
}
