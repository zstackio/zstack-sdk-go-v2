// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIpAddress deletes IpAddress
func (cli *ZSClient) DeleteIpAddress(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/ip-address", uuid, string(deleteMode))
}
