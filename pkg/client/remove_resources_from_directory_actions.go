// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveResourcesFromDirectory removes ResourcesFromDirectory
func (cli *ZSClient) RemoveResourcesFromDirectory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/remove/resources/directory", uuid, string(deleteMode))
}
