// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcHaGroup updates VpcHaGroup
func (cli *ZSClient) UpdateVpcHaGroup(uuid string, params param.UpdateVpcHaGroupParam) (*view.VpcHaGroupInventoryView, error) {
	var resp view.UpdateVpcHaGroupEventView
	if err := cli.Put("v1/vpc/hagroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVpcHaGroup deletes VpcHaGroup
func (cli *ZSClient) DeleteVpcHaGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/hagroups/{uuid}", uuid, string(deleteMode))
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
