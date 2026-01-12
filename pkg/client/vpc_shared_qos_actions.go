// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVpcSharedQos creates VpcSharedQos
func (cli *ZSClient) CreateVpcSharedQos(params param.CreateVpcSharedQosParam) (*view.VpcSharedQosInventoryView, error) {
	var resp view.CreateVpcSharedQosEventView
	if err := cli.Post("v1/vips/sharedqos", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVpcSharedQos queries VpcSharedQos list
func (cli *ZSClient) QueryVpcSharedQos(params *param.QueryParam) ([]view.VpcSharedQosInventoryView, error) {
	var resp []view.VpcSharedQosInventoryView
	return resp, cli.List("v1/vips/sharedqos", params, &resp)
}

func (cli *ZSClient) GetVpcSharedQos(uuid string) (*view.VpcSharedQosInventoryView, error) {
	var resp view.VpcSharedQosInventoryView
	if err := cli.Get("v1/vips/sharedqos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVpcSharedQos updates VpcSharedQos
func (cli *ZSClient) UpdateVpcSharedQos(uuid string, params param.UpdateVpcSharedQosParam) (*view.VpcSharedQosInventoryView, error) {
	var resp view.UpdateVpcSharedQosEventView
	err := cli.PutWithSpec("v1/vips/sharedqos", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVpcSharedQos deletes VpcSharedQos
func (cli *ZSClient) DeleteVpcSharedQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vips/sharedqos", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
