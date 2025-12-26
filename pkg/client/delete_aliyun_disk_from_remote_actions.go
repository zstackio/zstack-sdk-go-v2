// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunDiskFromRemote deletes AliyunDiskFromRemote
func (cli *ZSClient) DeleteAliyunDiskFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk/{uuid}/remote", uuid, string(deleteMode))
}
