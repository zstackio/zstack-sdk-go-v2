// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteContainerResourceFromEndpoint deletes ContainerResourceFromEndpoint
func (cli *ZSClient) DeleteContainerResourceFromEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/container/management/endpoint/{uuid}/resources/delete", uuid, string(deleteMode))
}
