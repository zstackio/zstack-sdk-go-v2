// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunKeySecret deletes AliyunKeySecret
func (cli *ZSClient) DeleteAliyunKeySecret(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/key/{uuid}", uuid, string(deleteMode))
}
