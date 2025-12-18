// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunDiskFromLocal deletes AliyunDiskFromLocal
func (cli *ZSClient) DeleteAliyunDiskFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk/{uuid}", uuid, string(deleteMode))
}
