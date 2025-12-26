// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteOssBucketRemote deletes OssBucketRemote
func (cli *ZSClient) DeleteOssBucketRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket/remote/{uuid}", uuid, string(deleteMode))
}
