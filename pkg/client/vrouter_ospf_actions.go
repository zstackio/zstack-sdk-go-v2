// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// UpdateVRouterOspfArea updates VRouterOspfArea
func (cli *ZSClient) UpdateVRouterOspfArea(uuid string, params param.UpdateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	resp := view.RouterAreaInventoryView{}
	if err := cli.PutWithRespKey("v1/routerArea", uuid, "", map[string]interface{}{
		"updateVRouterOspfArea": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVRouterOspfArea creates VRouterOspfArea
func (cli *ZSClient) CreateVRouterOspfArea(params param.CreateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	resp := view.RouterAreaInventoryView{}
	if err := cli.Post("v1/routerArea", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedOspf gets VpcAttachedOspf by uuid
func (cli *ZSClient) GetVpcAttachedOspf(uuid string, params param.GetVpcAttachedOspfParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-ospf", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromOspfArea removes VRouterNetworksFromOspfArea
func (cli *ZSClient) RemoveVRouterNetworksFromOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea/networks", uuid, string(deleteMode))
}

// SetVRouterRouterId operates on VRouterRouterId
func (cli *ZSClient) SetVRouterRouterId(vRouterUuid string, params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.Post(fmt.Sprintf("v1/routerArea/%s/routerid", vRouterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVRouterNetworksToOspfArea adds VRouterNetworksToOspfArea
func (cli *ZSClient) AddVRouterNetworksToOspfArea(routerAreaUuid string, vRouterUuid string, params param.AddVRouterNetworksToOspfAreaParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/routerArea/%s/router/%s/addnetworks", routerAreaUuid, vRouterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterRouterId gets VRouterRouterId by uuid
func (cli *ZSClient) GetVRouterRouterId(uuid string) (*view.GetVRouterRouterIdView, error) {
	var resp view.GetVRouterRouterIdView
	if err := cli.GetWithRespKey("v1/routerArea", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterOspfNeighbor gets VRouterOspfNeighbor by uuid
func (cli *ZSClient) GetVRouterOspfNeighbor(uuid string) (*view.GetVRouterOspfNeighborView, error) {
	var resp view.GetVRouterOspfNeighborView
	if err := cli.GetWithRespKey("v1/routerArea", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterOspfNetwork queries VRouterOspfNetwork list
func (cli *ZSClient) QueryVRouterOspfNetwork(params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, error) {
	var resp []view.NetworkRouterAreaRefInventoryView
	return resp, cli.List("v1/routerArea/network", params, &resp)
}

func (cli *ZSClient) GetVRouterOspfNetwork(uuid string) (*view.NetworkRouterAreaRefInventoryView, error) {
	var resp view.NetworkRouterAreaRefInventoryView
	if err := cli.Get("v1/routerArea/network", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterOspfNetwork Pagination
func (cli *ZSClient) PageVRouterOspfNetwork(params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, int, error) {
	var vRouterOspfNetworks []view.NetworkRouterAreaRefInventoryView
	total, err := cli.Page("v1/routerArea/network", params, &vRouterOspfNetworks)
	return vRouterOspfNetworks, total, err
}

// QueryVRouterOspfArea queries VRouterOspfArea list
func (cli *ZSClient) QueryVRouterOspfArea(params *param.QueryParam) ([]view.RouterAreaInventoryView, error) {
	var resp []view.RouterAreaInventoryView
	return resp, cli.List("v1/routerArea", params, &resp)
}

func (cli *ZSClient) GetVRouterOspfArea(uuid string) (*view.RouterAreaInventoryView, error) {
	var resp view.RouterAreaInventoryView
	if err := cli.Get("v1/routerArea", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterOspfArea Pagination
func (cli *ZSClient) PageVRouterOspfArea(params *param.QueryParam) ([]view.RouterAreaInventoryView, int, error) {
	var vRouterOspfAreas []view.RouterAreaInventoryView
	total, err := cli.Page("v1/routerArea", params, &vRouterOspfAreas)
	return vRouterOspfAreas, total, err
}

// DeleteVRouterOspfArea deletes VRouterOspfArea
func (cli *ZSClient) DeleteVRouterOspfArea(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/routerArea", uuid, string(deleteMode))
}
