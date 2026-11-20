// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

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

func (cli *ZSClient) GetMulticastRouter(uuid string) (*view.MulticastRouterInventoryView, error) {
	var resp view.MulticastRouterInventoryView
	if err := cli.Get("v1/multicast/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// RemoveRendezvousPointFromMulticastRouter requires additional arguments.
//
// Deprecated: use RemoveRendezvousPointFromMulticastRouterWithParams to supply the required fields.
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveRendezvousPointFromMulticastRouter requires additional arguments; use RemoveRendezvousPointFromMulticastRouterWithParams")
}

// RemoveRendezvousPointFromMulticastRouterWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouterWithParams(uuid string, deleteMode param.DeleteMode, params param.RemoveRendezvousPointFromMulticastRouterParamDetail) error {
	query := url.Values{
		"rpAddress":    {params.RpAddress},
		"groupAddress": {params.GroupAddress},
	}
	return cli.deleteActionWithQuery("v1/multicast/virtual-routers", uuid, "RendezvousPoint", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "rpAddress", "groupAddress")
}

// ChangeMulticastRouterState changes MulticastRouterState
func (cli *ZSClient) ChangeMulticastRouterState(uuid string, params param.ChangeMulticastRouterStateParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.PutWithRespKey("v1/multicast/virtual-routers", uuid, "", map[string]interface{}{
		"changeMulticastRouterState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcMulticastRoute gets VpcMulticastRoute by uuid
func (cli *ZSClient) GetVpcMulticastRoute(uuid string) (*view.MulticastRouteInventoryView, error) {
	var resp view.MulticastRouteInventoryView
	if err := cli.GetWithRespKey("v1/multicast/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRendezvousPointToMulticastRouter adds RendezvousPointToMulticastRouter
func (cli *ZSClient) AddRendezvousPointToMulticastRouter(uuid string, params param.AddRendezvousPointToMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/multicast/virtual-routers/%s/RendezvousPoint", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
