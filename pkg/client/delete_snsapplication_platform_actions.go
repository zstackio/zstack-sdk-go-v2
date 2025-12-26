// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSNSApplicationPlatform deletes SNSApplicationPlatform
func (cli *ZSClient) DeleteSNSApplicationPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-platforms/{uuid}", uuid, string(deleteMode))
}
