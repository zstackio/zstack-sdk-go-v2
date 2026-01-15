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
	resp := view.VpcHaGroupInventoryView{}
	if err := cli.Put("v1/vpc/hagroups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVpcHaGroup deletes VpcHaGroup
func (cli *ZSClient) DeleteVpcHaGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpc/hagroups", uuid, string(deleteMode))
}
// CreateVpcHaGroup creates VpcHaGroup
func (cli *ZSClient) CreateVpcHaGroup(params param.CreateVpcHaGroupParam) (*view.VpcHaGroupInventoryView, error) {
	resp := view.VpcHaGroupInventoryView{}
	if err := cli.Post("v1/vpc/hagroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVpcHaGroup queries VpcHaGroup list
func (cli *ZSClient) QueryVpcHaGroup(params *param.QueryParam) ([]view.VpcHaGroupInventoryView, error) {
	var resp []view.VpcHaGroupInventoryView
	return resp, cli.List("v1/vpc/hagroups", params, &resp)
}

// PageVpcHaGroup Pagination
func (cli *ZSClient) PageVpcHaGroup(params *param.QueryParam) ([]view.VpcHaGroupInventoryView, int, error) {
	var vpcHaGroups []view.VpcHaGroupInventoryView
	total, err := cli.Page("v1/vpc/hagroups", params, &vpcHaGroups)
	return vpcHaGroups, total, err
}
