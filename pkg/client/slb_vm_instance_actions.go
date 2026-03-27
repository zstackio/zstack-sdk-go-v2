// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySlbVmInstance queries SlbVmInstance list
func (cli *ZSClient) QuerySlbVmInstance(ctx context.Context, params *param.QueryParam) ([]view.SlbVmInstanceInventoryView, error) {
	var resp []view.SlbVmInstanceInventoryView
	return resp, cli.List(ctx, "v1/load-balancers/slb/instances", params, &resp)
}

func (cli *ZSClient) GetSlbVmInstance(ctx context.Context, uuid string) (*view.SlbVmInstanceInventoryView, error) {
	var resp view.SlbVmInstanceInventoryView
	if err := cli.Get(ctx, "v1/load-balancers/slb/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSlbVmInstance Pagination
func (cli *ZSClient) PageSlbVmInstance(ctx context.Context, params *param.QueryParam) ([]view.SlbVmInstanceInventoryView, int, error) {
	var slbVmInstances []view.SlbVmInstanceInventoryView
	total, err := cli.Page(ctx, "v1/load-balancers/slb/instances", params, &slbVmInstances)
	return slbVmInstances, total, err
}
