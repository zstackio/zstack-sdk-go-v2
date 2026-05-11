// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryOvnController queries OvnController list
func (cli *ZSClient) QueryOvnController(ctx context.Context, params *param.QueryParam) ([]view.OvnControllerInventoryView, error) {
	var resp []view.OvnControllerInventoryView
	return resp, cli.List(ctx, "v1/ovn-controllers", params, &resp)
}

func (cli *ZSClient) GetOvnController(ctx context.Context, uuid string) (*view.OvnControllerInventoryView, error) {
	var resp view.OvnControllerInventoryView
	if err := cli.Get(ctx, "v1/ovn-controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageOvnController Pagination
func (cli *ZSClient) PageOvnController(ctx context.Context, params *param.QueryParam) ([]view.OvnControllerInventoryView, int, error) {
	var ovnControllers []view.OvnControllerInventoryView
	total, err := cli.Page(ctx, "v1/ovn-controllers", params, &ovnControllers)
	return ovnControllers, total, err
}
