// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHuaweiIMasterTenant queries HuaweiIMasterTenant list
func (cli *ZSClient) QueryHuaweiIMasterTenant(params *param.QueryParam) ([]view.HuaweiIMasterTenantInventoryView, error) {
	var resp []view.HuaweiIMasterTenantInventoryView
	return resp, cli.List("v1/sdn-controller/huawei-imaster/tenants", params, &resp)
}
