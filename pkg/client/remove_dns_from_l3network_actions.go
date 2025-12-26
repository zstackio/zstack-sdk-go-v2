// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveDnsFromL3Network removes DnsFromL3Network
func (cli *ZSClient) RemoveDnsFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/dns/{dns}", uuid, string(deleteMode))
}
