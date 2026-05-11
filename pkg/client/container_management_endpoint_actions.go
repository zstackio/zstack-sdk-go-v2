// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateContainerManagementEndpoint updates ContainerManagementEndpoint
func (cli *ZSClient) UpdateContainerManagementEndpoint(ctx context.Context, uuid string, params param.UpdateContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	resp := view.ContainerManagementEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/container/management/endpoint", uuid, "", map[string]interface{}{
		"updateContainerManagementEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryContainerManagementEndpoint queries ContainerManagementEndpoint list
func (cli *ZSClient) QueryContainerManagementEndpoint(ctx context.Context, params *param.QueryParam) ([]view.ContainerManagementEndpointInventoryView, error) {
	var resp []view.ContainerManagementEndpointInventoryView
	return resp, cli.List(ctx, "v1/container/management/endpoint", params, &resp)
}

func (cli *ZSClient) GetContainerManagementEndpoint(ctx context.Context, uuid string) (*view.ContainerManagementEndpointInventoryView, error) {
	var resp view.ContainerManagementEndpointInventoryView
	if err := cli.Get(ctx, "v1/container/management/endpoint", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageContainerManagementEndpoint Pagination
func (cli *ZSClient) PageContainerManagementEndpoint(ctx context.Context, params *param.QueryParam) ([]view.ContainerManagementEndpointInventoryView, int, error) {
	var containerManagementEndpoints []view.ContainerManagementEndpointInventoryView
	total, err := cli.Page(ctx, "v1/container/management/endpoint", params, &containerManagementEndpoints)
	return containerManagementEndpoints, total, err
}
// AddContainerManagementEndpoint adds ContainerManagementEndpoint
func (cli *ZSClient) AddContainerManagementEndpoint(ctx context.Context, params param.AddContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	resp := view.ContainerManagementEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/container/management/endpoint", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SyncContainerManagementEndpoint operates on ContainerManagementEndpoint
func (cli *ZSClient) SyncContainerManagementEndpoint(ctx context.Context, uuid string, params param.SyncContainerManagementEndpointParam) (*view.ContainerManagementEndpointInventoryView, error) {
	resp := view.ContainerManagementEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/container/management/endpoint", uuid, "", map[string]interface{}{
		"syncContainerManagementEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteContainerManagementEndpoint deletes ContainerManagementEndpoint
func (cli *ZSClient) DeleteContainerManagementEndpoint(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/container/management/endpoint", uuid, string(deleteMode))
}
