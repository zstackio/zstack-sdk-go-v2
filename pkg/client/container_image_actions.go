// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryContainerImage queries ContainerImage list
func (cli *ZSClient) QueryContainerImage(ctx context.Context, params *param.QueryParam) ([]view.ContainerImageInventoryView, error) {
	var resp []view.ContainerImageInventoryView
	return resp, cli.List(ctx, "v1/container/images", params, &resp)
}

func (cli *ZSClient) GetContainerImage(ctx context.Context, uuid string) (*view.ContainerImageInventoryView, error) {
	var resp view.ContainerImageInventoryView
	if err := cli.Get(ctx, "v1/container/images", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageContainerImage Pagination
func (cli *ZSClient) PageContainerImage(ctx context.Context, params *param.QueryParam) ([]view.ContainerImageInventoryView, int, error) {
	var containerImages []view.ContainerImageInventoryView
	total, err := cli.Page(ctx, "v1/container/images", params, &containerImages)
	return containerImages, total, err
}
