// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHuaweiIMasterFabric deletes HuaweiIMasterFabric
func (cli *ZSClient) DeleteHuaweiIMasterFabric(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/fabrics/{uuid}", uuid, string(deleteMode))
}
