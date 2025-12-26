// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAppBuildSystem deletes AppBuildSystem
func (cli *ZSClient) DeleteAppBuildSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildsystem/{uuid}", uuid, string(deleteMode))
}
