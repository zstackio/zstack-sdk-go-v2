// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryApplianceVm queries ApplianceVm list
func (cli *ZSClient) QueryApplianceVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List("v1/vm-instances/appliances", params, &resp)
}
