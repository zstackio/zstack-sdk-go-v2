// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteL3Network deletes L3Network
func (cli *ZSClient) DeleteL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{uuid}", uuid, string(deleteMode))
}
