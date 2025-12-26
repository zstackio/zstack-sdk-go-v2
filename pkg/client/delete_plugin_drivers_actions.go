// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePluginDrivers deletes PluginDrivers
func (cli *ZSClient) DeletePluginDrivers(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/external/plugins/{uuid}", uuid, string(deleteMode))
}
