// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateMulticastRouter creates MulticastRouter
func (cli *ZSClient) CreateMulticastRouter(params param.CreateMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	var resp view.CreateMulticastRouterEventView
	if err := cli.Post("v1/multicast/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryMulticastRouter queries MulticastRouter list
func (cli *ZSClient) QueryMulticastRouter(params *param.QueryParam) ([]view.MulticastRouterInventoryView, error) {
	var resp []view.MulticastRouterInventoryView
	return resp, cli.List("v1/multicast/virtual-routers", params, &resp)
}
// DeleteMulticastRouter deletes MulticastRouter
func (cli *ZSClient) DeleteMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/multicast/virtual-routers/{uuid}", uuid, string(deleteMode))
}
