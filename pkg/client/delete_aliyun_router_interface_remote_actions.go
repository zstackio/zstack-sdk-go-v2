// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunRouterInterfaceRemote deletes AliyunRouterInterfaceRemote
func (cli *ZSClient) DeleteAliyunRouterInterfaceRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/remote/{uuid}", uuid, string(deleteMode))
}
