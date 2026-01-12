// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateEip creates Eip
func (cli *ZSClient) CreateEip(params param.CreateEipParam) (*view.EipInventoryView, error) {
	var resp view.CreateEipEventView
	if err := cli.Post("v1/eips", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AttachEip operates on Eip
func (cli *ZSClient) AttachEip(eipUuid string, vmNicUuid string, params param.AttachEipParam) (*view.EipInventoryView, error) {
	var resp view.AttachEipEventView
	err := cli.PostWithSpec("v1/eips", fmt.Sprintf(\"%s/vm-instances/nics/%s\", eipUuid, vmNicUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateEip updates Eip
func (cli *ZSClient) UpdateEip(uuid string, params param.UpdateEipParam) (*view.EipInventoryView, error) {
	var resp view.UpdateEipEventView
	if err := cli.Put("v1/eips", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryEip queries Eip list
func (cli *ZSClient) QueryEip(params *param.QueryParam) ([]view.EipInventoryView, error) {
	var resp []view.EipInventoryView
	return resp, cli.List("v1/eips", params, &resp)
}

func (cli *ZSClient) GetEip(uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.Get("v1/eips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteEip deletes Eip
func (cli *ZSClient) DeleteEip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/eips", uuid, string(deleteMode))
}
// DetachEip operates on Eip
func (cli *ZSClient) DetachEip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/eips", uuid, string(deleteMode))
}
