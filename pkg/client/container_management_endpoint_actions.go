// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateContainerManagementEndpoint updates ContainerManagementEndpoint
func (cli *ZSClient) UpdateContainerManagementEndpoint(uuid string, params param.UpdateContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	var resp view.UpdateContainerManagementEndpointEventView
	if err := cli.Put("v1/container/management/endpoint", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryContainerManagementEndpoint queries ContainerManagementEndpoint list
func (cli *ZSClient) QueryContainerManagementEndpoint(params *param.QueryParam) ([]view.ContainerManagementEndpointInventoryView, error) {
	var resp []view.ContainerManagementEndpointInventoryView
	return resp, cli.List("v1/container/management/endpoint", params, &resp)
}

func (cli *ZSClient) GetContainerManagementEndpoint(uuid string) (*view.ContainerManagementEndpointInventoryView, error) {
	var resp view.ContainerManagementEndpointInventoryView
	if err := cli.Get("v1/container/management/endpoint", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddContainerManagementEndpoint adds ContainerManagementEndpoint
func (cli *ZSClient) AddContainerManagementEndpoint(params param.AddContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	var resp view.AddContainerManagementEndpointEventView
	if err := cli.Post("v1/container/management/endpoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// SyncContainerManagementEndpoint operates on ContainerManagementEndpoint
func (cli *ZSClient) SyncContainerManagementEndpoint(uuid string, params param.SyncContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	var resp view.SyncContainerManagementEndpointEventView
	if err := cli.Put("v1/container/management/endpoint", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteContainerManagementEndpoint deletes ContainerManagementEndpoint
func (cli *ZSClient) DeleteContainerManagementEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/container/management/endpoint", uuid, string(deleteMode))
}
