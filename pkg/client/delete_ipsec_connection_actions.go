// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIPsecConnection deletes IPsecConnection
func (cli *ZSClient) DeleteIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec/{uuid}", uuid, string(deleteMode))
}
