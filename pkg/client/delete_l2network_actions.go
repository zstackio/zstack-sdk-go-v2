// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteL2Network deletes L2Network
func (cli *ZSClient) DeleteL2Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l2-networks/{uuid}", uuid, string(deleteMode))
}
