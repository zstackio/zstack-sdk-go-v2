// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHuaweiIMasterVpc queries HuaweiIMasterVpc list
func (cli *ZSClient) QueryHuaweiIMasterVpc(params *param.QueryParam) ([]view.HuaweiIMasterVpcInventoryView, error) {
	var resp []view.HuaweiIMasterVpcInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/vpcs", params, &resp)
}
