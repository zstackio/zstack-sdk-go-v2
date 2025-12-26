// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteNicQos deletes NicQos
func (cli *ZSClient) DeleteNicQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/nic-qos", uuid, string(deleteMode))
}
