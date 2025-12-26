// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHuaweiIMasterTenant deletes HuaweiIMasterTenant
func (cli *ZSClient) DeleteHuaweiIMasterTenant(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/tenants/{uuid}", uuid, string(deleteMode))
}
