// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachNetworkServiceFromL3Network operates on NetworkServiceFromL3Network
func (cli *ZSClient) DetachNetworkServiceFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/{l3NetworkUuid}/network-services", uuid, string(deleteMode))
}
