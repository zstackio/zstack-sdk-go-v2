// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHuaweiIMasterTenant 删除HuaweiIMasterTenant
func (cli *ZSClient) DeleteHuaweiIMasterTenant(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/tenants/{uuid}", uuid, string(deleteMode))
}

