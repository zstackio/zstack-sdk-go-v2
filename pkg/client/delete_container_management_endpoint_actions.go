// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteContainerManagementEndpoint deletes ContainerManagementEndpoint
func (cli *ZSClient) DeleteContainerManagementEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/container/management/endpoint/{uuid}", uuid, string(deleteMode))
}
