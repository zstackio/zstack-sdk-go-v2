// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePublishApp deletes PublishApp
func (cli *ZSClient) DeletePublishApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/app/{uuid}", uuid, string(deleteMode))
}
