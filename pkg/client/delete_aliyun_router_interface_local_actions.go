// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunRouterInterfaceLocal deletes AliyunRouterInterfaceLocal
func (cli *ZSClient) DeleteAliyunRouterInterfaceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/{uuid}", uuid, string(deleteMode))
}
