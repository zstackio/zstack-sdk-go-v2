// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteIpRange deletes IpRange
func (cli *ZSClient) DeleteIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/ip-ranges/{uuid}", uuid, string(deleteMode))
}
