// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteOssBucketFileRemote 删除OssBucketFileRemote
func (cli *ZSClient) DeleteOssBucketFileRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket-file/remote/{uuid}", uuid, string(deleteMode))
}

