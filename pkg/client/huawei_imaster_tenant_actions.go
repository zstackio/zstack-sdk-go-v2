// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteHuaweiIMasterTenant deletes HuaweiIMasterTenant
func (cli *ZSClient) DeleteHuaweiIMasterTenant(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/tenants", uuid, string(deleteMode))
}
// QueryHuaweiIMasterTenant queries HuaweiIMasterTenant list
func (cli *ZSClient) QueryHuaweiIMasterTenant(params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, error) {
	var resp []view.HuaweiIMasterTenantInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/tenants", params, &resp)
}

// PageHuaweiIMasterTenant Pagination
func (cli *ZSClient) PageHuaweiIMasterTenant(params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, int, error) {
	var huaweiIMasterTenants []view.HuaweiIMasterTenantInventoryView
	total, err := cli.Page("v1/sdn-controller/huawei-imaster/tenants", params, &huaweiIMasterTenants)
	return huaweiIMasterTenants, total, err
}
