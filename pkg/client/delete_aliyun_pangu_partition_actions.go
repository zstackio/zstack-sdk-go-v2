// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunPanguPartition deletes AliyunPanguPartition
func (cli *ZSClient) DeleteAliyunPanguPartition(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun/pangu/{uuid}", uuid, string(deleteMode))
}
