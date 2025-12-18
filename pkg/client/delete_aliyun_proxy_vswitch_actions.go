// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAliyunProxyVSwitch deletes AliyunProxyVSwitch
func (cli *ZSClient) DeleteAliyunProxyVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun-proxy/vpcs/vswitches/{uuid}", uuid, string(deleteMode))
}
