// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteOssBucketFileRemote deletes OssBucketFileRemote
func (cli *ZSClient) DeleteOssBucketFileRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket-file/remote/{uuid}", uuid, string(deleteMode))
}
