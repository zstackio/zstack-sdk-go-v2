// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcHaGroup updates VpcHaGroup
func (cli *ZSClient) UpdateVpcHaGroup(uuid string, params param.UpdateVpcHaGroupParam) (*view.VpcHaGroupInventoryView, error) {
	var resp view.UpdateVpcHaGroupEventView
	err := cli.PutWithSpec("v1/vpc/hagroups", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVpcHaGroup deletes VpcHaGroup
func (cli *ZSClient) DeleteVpcHaGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vpc/hagroups", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// CreateVpcHaGroup creates VpcHaGroup
func (cli *ZSClient) CreateVpcHaGroup(params param.CreateVpcHaGroupParam) (*view.VpcHaGroupInventoryView, error) {
	var resp view.CreateVpcHaGroupEventView
	if err := cli.Post("v1/vpc/hagroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVpcHaGroup queries VpcHaGroup list
func (cli *ZSClient) QueryVpcHaGroup(params *param.QueryParam) ([]view.VpcHaGroupInventoryView, error) {
	var resp []view.VpcHaGroupInventoryView
	return resp, cli.List("v1/vpc/hagroups", params, &resp)
}

func (cli *ZSClient) GetVpcHaGroup(uuid string) (*view.VpcHaGroupInventoryView, error) {
	var resp view.VpcHaGroupInventoryView
	if err := cli.Get("v1/vpc/hagroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
