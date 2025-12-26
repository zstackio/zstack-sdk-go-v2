// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunProxyVSwitch deletes AliyunProxyVSwitch
func (cli *ZSClient) DeleteAliyunProxyVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun-proxy/vpcs/vswitches/{uuid}", uuid, string(deleteMode))
}
