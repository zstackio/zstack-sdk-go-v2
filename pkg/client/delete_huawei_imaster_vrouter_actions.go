// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHuaweiIMasterVRouter deletes HuaweiIMasterVRouter
func (cli *ZSClient) DeleteHuaweiIMasterVRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/vrouters/{uuid}", uuid, string(deleteMode))
}
