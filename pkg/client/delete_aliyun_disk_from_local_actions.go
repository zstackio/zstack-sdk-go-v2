// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunDiskFromLocal deletes AliyunDiskFromLocal
func (cli *ZSClient) DeleteAliyunDiskFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk/{uuid}", uuid, string(deleteMode))
}
