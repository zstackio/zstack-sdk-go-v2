// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVniRange updates VniRange
func (cli *ZSClient) UpdateVniRange(uuid string, params param.UpdateVniRangeParam) (*view.VniRangeInventoryView, error) {
	var resp view.UpdateVniRangeEventView
	if err := cli.Put("v1/l2-networks/vxlan-pool/vni-ranges/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVniRange queries VniRange list
func (cli *ZSClient) QueryVniRange(params *param.QueryParam) ([]view.VniRangeInventoryView, error) {
	var resp []view.VniRangeInventoryView
	return resp, cli.List("v1/l2-networks/vxlan-pool/vni-range", params, &resp)
}
// CreateVniRange creates VniRange
func (cli *ZSClient) CreateVniRange(params param.CreateVniRangeParam) (*view.VniRangeInventoryView, error) {
	var resp view.CreateVniRangeEventView
	if err := cli.Post("v1/l2-networks/vxlan-pool/{l2NetworkUuid}/vni-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVniRange deletes VniRange
func (cli *ZSClient) DeleteVniRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/vxlan-pool/vni-ranges/{uuid}", uuid, string(deleteMode))
}
