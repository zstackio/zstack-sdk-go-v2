// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunSnapshotFromRemote deletes AliyunSnapshotFromRemote
func (cli *ZSClient) DeleteAliyunSnapshotFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot/{uuid}/remote", uuid, string(deleteMode))
}
