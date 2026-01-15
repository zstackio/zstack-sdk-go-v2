// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateMulticastRouter creates MulticastRouter
func (cli *ZSClient) CreateMulticastRouter(params param.CreateMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.Post("v1/multicast/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMulticastRouter queries MulticastRouter list
func (cli *ZSClient) QueryMulticastRouter(params *param.QueryParam) ([]view.MulticastRouterInventoryView, error) {
	var resp []view.MulticastRouterInventoryView
	return resp, cli.List("v1/multicast/virtual-routers", params, &resp)
}

// PageMulticastRouter Pagination
func (cli *ZSClient) PageMulticastRouter(params *param.QueryParam) ([]view.MulticastRouterInventoryView, int, error) {
	var multicastRouters []view.MulticastRouterInventoryView
	total, err := cli.Page("v1/multicast/virtual-routers", params, &multicastRouters)
	return multicastRouters, total, err
}
// DeleteMulticastRouter deletes MulticastRouter
func (cli *ZSClient) DeleteMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/multicast/virtual-routers", uuid, string(deleteMode))
}
