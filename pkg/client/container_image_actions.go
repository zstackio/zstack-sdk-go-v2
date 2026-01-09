// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryContainerImage queries ContainerImage list
func (cli *ZSClient) QueryContainerImage(params *param.QueryParam) ([]view.ContainerImageInventoryView, error) {
	var resp []view.ContainerImageInventoryView
	return resp, cli.List("v1/container/images", params, &resp)
}

func (cli *ZSClient) GetContainerImage(uuid string) (*view.ContainerImageInventoryView, error) {
	var resp view.ContainerImageInventoryView
	if err := cli.Get("v1/container/images", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
