// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunRouterInterfaceLocal 删除AliyunRouterInterfaceLocal
func (cli *ZSClient) DeleteAliyunRouterInterfaceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/{uuid}", uuid, string(deleteMode))
}

