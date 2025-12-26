// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteInstanceOffering deletes InstanceOffering
func (cli *ZSClient) DeleteInstanceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/instance-offerings/{uuid}", uuid, string(deleteMode))
}
