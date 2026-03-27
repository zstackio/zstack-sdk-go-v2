// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteHuaweiIMasterTenant deletes HuaweiIMasterTenant
func (cli *ZSClient) DeleteHuaweiIMasterTenant(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sdn-controller/huawei-imaster/tenants", uuid, string(deleteMode))
}
// QueryHuaweiIMasterTenant queries HuaweiIMasterTenant list
func (cli *ZSClient) QueryHuaweiIMasterTenant(ctx context.Context, params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, error) {
	var resp []view.HuaweiIMasterTenantInventoryView
	return resp, cli.List(ctx, "v1/sdn-controller/huawei-imaster/tenants", params, &resp)
}

func (cli *ZSClient) GetHuaweiIMasterTenant(ctx context.Context, uuid string) (*view.HuaweiIMasterTenantInventoryView, error) {
	var resp view.HuaweiIMasterTenantInventoryView
	if err := cli.Get(ctx, "v1/sdn-controller/huawei-imaster/tenants", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHuaweiIMasterTenant Pagination
func (cli *ZSClient) PageHuaweiIMasterTenant(ctx context.Context, params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, int, error) {
	var huaweiIMasterTenants []view.HuaweiIMasterTenantInventoryView
	total, err := cli.Page(ctx, "v1/sdn-controller/huawei-imaster/tenants", params, &huaweiIMasterTenants)
	return huaweiIMasterTenants, total, err
}
