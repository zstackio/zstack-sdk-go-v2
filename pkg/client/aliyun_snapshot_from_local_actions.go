// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunSnapshotFromLocal 删除AliyunSnapshotFromLocal
func (cli *ZSClient) DeleteAliyunSnapshotFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot/{uuid}", uuid, string(deleteMode))
}

