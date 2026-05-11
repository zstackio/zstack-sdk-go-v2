// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryZnsTenant queries ZnsTenant list
func (cli *ZSClient) QueryZnsTenant(ctx context.Context, params *param.QueryParam) ([]view.ZnsTenantInventoryView, error) {
	var resp []view.ZnsTenantInventoryView
	return resp, cli.List(ctx, "v1/sdn-controller/zns/tenants", params, &resp)
}

func (cli *ZSClient) GetZnsTenant(ctx context.Context, uuid string) (*view.ZnsTenantInventoryView, error) {
	var resp view.ZnsTenantInventoryView
	if err := cli.Get(ctx, "v1/sdn-controller/zns/tenants", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZnsTenant Pagination
func (cli *ZSClient) PageZnsTenant(ctx context.Context, params *param.QueryParam) ([]view.ZnsTenantInventoryView, int, error) {
	var znsTenants []view.ZnsTenantInventoryView
	total, err := cli.Page(ctx, "v1/sdn-controller/zns/tenants", params, &znsTenants)
	return znsTenants, total, err
}
