// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVmNic creates VmNic
func (cli *ZSClient) CreateVmNic(params param.CreateVmNicParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.Post("v1/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmNic queries VmNic list
func (cli *ZSClient) QueryVmNic(params *param.QueryParam) ([]view.VmNicInventoryView, error) {
	var resp []view.VmNicInventoryView
	return resp, cli.List("v1/vm-instances/nics", params, &resp)
}

func (cli *ZSClient) GetVmNic(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/vm-instances/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmNic Pagination
func (cli *ZSClient) PageVmNic(params *param.QueryParam) ([]view.VmNicInventoryView, int, error) {
	var vmNics []view.VmNicInventoryView
	total, err := cli.Page("v1/vm-instances/nics", params, &vmNics)
	return vmNics, total, err
}
// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nics", uuid, string(deleteMode))
}
