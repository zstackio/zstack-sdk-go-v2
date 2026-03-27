// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryOvnControllerVmInstance queries OvnControllerVmInstance list
func (cli *ZSClient) QueryOvnControllerVmInstance(ctx context.Context, params *param.QueryParam) ([]view.OvnControllerVmInstanceInventoryView, error) {
	var resp []view.OvnControllerVmInstanceInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/appliances/ovn-controller", params, &resp)
}

func (cli *ZSClient) GetOvnControllerVmInstance(ctx context.Context, uuid string) (*view.OvnControllerVmInstanceInventoryView, error) {
	var resp view.OvnControllerVmInstanceInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/appliances/ovn-controller", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageOvnControllerVmInstance Pagination
func (cli *ZSClient) PageOvnControllerVmInstance(ctx context.Context, params *param.QueryParam) ([]view.OvnControllerVmInstanceInventoryView, int, error) {
	var ovnControllerVmInstances []view.OvnControllerVmInstanceInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/appliances/ovn-controller", params, &ovnControllerVmInstances)
	return ovnControllerVmInstances, total, err
}
