// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunRouteEntryRemote deletes AliyunRouteEntryRemote
func (cli *ZSClient) DeleteAliyunRouteEntryRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/route-entry/{uuid}", uuid, string(deleteMode))
}
