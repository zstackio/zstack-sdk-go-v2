// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAliyunProxyVpc deletes AliyunProxyVpc
func (cli *ZSClient) DeleteAliyunProxyVpc(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun-proxy/vpcs/{uuid}", uuid, string(deleteMode))
}
