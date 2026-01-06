// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVmNic creates VmNic
func (cli *ZSClient) CreateVmNic(params param.CreateVmNicParam) (*view.VmNicInventoryView, error) {
	var resp view.CreateVmNicEventView
	if err := cli.Post("v1/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVmNic queries VmNic list
func (cli *ZSClient) QueryVmNic(params *param.QueryParam) ([]view.VmNicInventoryView, error) {
	var resp []view.VmNicInventoryView
	return resp, cli.List("v1/vm-instances/nics", params, &resp)
}
// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nics/{uuid}", uuid, string(deleteMode))
}
