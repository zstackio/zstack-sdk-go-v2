// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunDiskFromRemote deletes AliyunDiskFromRemote
func (cli *ZSClient) DeleteAliyunDiskFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk/{uuid}/remote", uuid, string(deleteMode))
}
