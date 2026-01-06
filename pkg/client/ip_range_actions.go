// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddIpRange adds IpRange
func (cli *ZSClient) AddIpRange(params param.AddIpRangeParam) (*view.IpRangeInventoryView, error) {
	var resp view.AddIpRangeEventView
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateIpRange updates IpRange
func (cli *ZSClient) UpdateIpRange(uuid string, params param.UpdateIpRangeParam) (*view.IpRangeInventoryView, error) {
	var resp view.UpdateIpRangeEventView
	if err := cli.Put("v1/l3-networks/ip-ranges/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIpRange queries IpRange list
func (cli *ZSClient) QueryIpRange(params *param.QueryParam) ([]view.IpRangeInventoryView, error) {
	var resp []view.IpRangeInventoryView
	return resp, cli.List("v1/l3-networks/ip-ranges", params, &resp)
}
// DeleteIpRange deletes IpRange
func (cli *ZSClient) DeleteIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/ip-ranges/{uuid}", uuid, string(deleteMode))
}
