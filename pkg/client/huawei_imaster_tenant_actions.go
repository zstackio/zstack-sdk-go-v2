// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteHuaweiIMasterTenant sends the action with its default options.
func (cli *ZSClient) DeleteHuaweiIMasterTenant(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteHuaweiIMasterTenantWithParams(uuid, deleteMode, param.DeleteHuaweiIMasterTenantParamDetail{})
}

// DeleteHuaweiIMasterTenantWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteHuaweiIMasterTenantWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteHuaweiIMasterTenantParamDetail) error {
	query := url.Values{}
	if params.SdnControllerUuid != nil {
		query.Set("sdnControllerUuid", *params.SdnControllerUuid)
	}
	return cli.deleteResourceWithQuery("v1/sdn-controller/huawei-imaster/tenants", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}
// QueryHuaweiIMasterTenant queries HuaweiIMasterTenant list
func (cli *ZSClient) QueryHuaweiIMasterTenant(params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, error) {
	var resp []view.HuaweiIMasterTenantInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/tenants", params, &resp)
}

func (cli *ZSClient) GetHuaweiIMasterTenant(uuid string) (*view.HuaweiIMasterTenantInventoryView, error) {
	var resp view.HuaweiIMasterTenantInventoryView
	if err := cli.Get("v1/sdn-controller/huawei-imaster/tenants", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHuaweiIMasterTenant Pagination
func (cli *ZSClient) PageHuaweiIMasterTenant(params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, int, error) {
	var huaweiIMasterTenants []view.HuaweiIMasterTenantInventoryView
	total, err := cli.Page("v1/sdn-controller/huawei-imaster/tenants", params, &huaweiIMasterTenants)
	return huaweiIMasterTenants, total, err
}
