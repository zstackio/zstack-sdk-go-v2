// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryOvnControllerVmInstance queries OvnControllerVmInstance list
func (cli *ZSClient) QueryOvnControllerVmInstance(params *param.QueryParam) ([]view.OvnControllerVmInstanceInventoryView, error) {
	var resp []view.OvnControllerVmInstanceInventoryView
	return resp, cli.List("v1/vm-instances/appliances/ovn-controller", params, &resp)
}
