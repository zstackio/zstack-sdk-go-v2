// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVmCdRom creates VmCdRom
func (cli *ZSClient) CreateVmCdRom(params param.CreateVmCdRomParam) (*view.VmCdRomInventoryView, error) {
	var resp view.CreateVmCdRomEventView
	if err := cli.Post("v1/vm-instances/cdroms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVmCdRom deletes VmCdRom
func (cli *ZSClient) DeleteVmCdRom(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/cdroms", uuid, string(deleteMode))
}
// QueryVmCdRom queries VmCdRom list
func (cli *ZSClient) QueryVmCdRom(params *param.QueryParam) ([]view.VmCdRomInventoryView, error) {
	var resp []view.VmCdRomInventoryView
	return resp, cli.List("v1/vm-instances/cdroms", params, &resp)
}

func (cli *ZSClient) GetVmCdRom(uuid string) (*view.VmCdRomInventoryView, error) {
	var resp view.VmCdRomInventoryView
	if err := cli.Get("v1/vm-instances/cdroms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmCdRom updates VmCdRom
func (cli *ZSClient) UpdateVmCdRom(uuid string, params param.UpdateVmCdRomParam) (*view.VmCdRomInventoryView, error) {
	var resp view.UpdateVmCdRomEventView
	if err := cli.Put("v1/vm-instances/cdroms", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
