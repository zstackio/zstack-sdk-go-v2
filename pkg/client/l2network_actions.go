// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteL2Network deletes L2Network
func (cli *ZSClient) DeleteL2Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{uuid}", uuid, string(deleteMode))
}
// UpdateL2Network updates L2Network
func (cli *ZSClient) UpdateL2Network(uuid string, params param.UpdateL2NetworkParam) (*view.L2NetworkInventoryView, error) {
	var resp view.UpdateL2NetworkEventView
	if err := cli.Put("v1/l2-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryL2Network queries L2Network list
func (cli *ZSClient) QueryL2Network(params *param.QueryParam) ([]view.L2NetworkInventoryView, error) {
	var resp []view.L2NetworkInventoryView
	return resp, cli.List("v1/l2-networks", params, &resp)
}
