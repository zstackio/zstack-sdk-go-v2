// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteOssBucketNameLocal deletes OssBucketNameLocal
func (cli *ZSClient) DeleteOssBucketNameLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket/{uuid}", uuid, string(deleteMode))
}
