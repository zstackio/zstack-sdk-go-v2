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

func (cli *ZSClient) GetVmNic(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.Get("v1/vm-instances/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVmNic deletes VmNic
func (cli *ZSClient) DeleteVmNic(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/nics", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
