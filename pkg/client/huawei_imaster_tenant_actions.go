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
	return cli.DeleteWithSpec("v1/sdn-controller/huawei-imaster/tenants", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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
