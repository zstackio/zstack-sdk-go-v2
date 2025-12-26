// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteHostNetworkServiceType deletes HostNetworkServiceType
func (cli *ZSClient) DeleteHostNetworkServiceType(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/service-types/{uuid}", uuid, string(deleteMode))
}
