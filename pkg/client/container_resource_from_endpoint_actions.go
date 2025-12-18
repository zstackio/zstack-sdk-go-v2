// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteContainerResourceFromEndpoint 删除ContainerResourceFromEndpoint
func (cli *ZSClient) DeleteContainerResourceFromEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/container/management/endpoint/{uuid}/resources/delete", uuid, string(deleteMode))
}

