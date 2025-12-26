// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBonding deletes Bonding
func (cli *ZSClient) DeleteBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/bondings/{uuid}", uuid, string(deleteMode))
}
