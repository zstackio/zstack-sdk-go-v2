// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHuaweiIMasterFabric queries HuaweiIMasterFabric list
func (cli *ZSClient) QueryHuaweiIMasterFabric(params *param.QueryParam) ([]view.HuaweiIMasterFabricInventoryView, error) {
	var resp []view.HuaweiIMasterFabricInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/fabrics", params, &resp)
}
