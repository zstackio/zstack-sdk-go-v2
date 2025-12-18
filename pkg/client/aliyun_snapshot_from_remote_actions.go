// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunSnapshotFromRemote 删除AliyunSnapshotFromRemote
func (cli *ZSClient) DeleteAliyunSnapshotFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot/{uuid}/remote", uuid, string(deleteMode))
}

