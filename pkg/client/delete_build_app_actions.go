// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBuildApp deletes BuildApp
func (cli *ZSClient) DeleteBuildApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildapp/{uuid}", uuid, string(deleteMode))
}
