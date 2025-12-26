// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteReservedIpRange deletes ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/reserved-ip-ranges/{uuid}", uuid, string(deleteMode))
}
