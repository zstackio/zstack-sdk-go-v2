// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePluginDrivers 删除PluginDrivers
func (cli *ZSClient) DeletePluginDrivers(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/external/plugins/{uuid}", uuid, string(deleteMode))
}

