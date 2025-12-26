// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSNSTopic deletes SNSTopic
func (cli *ZSClient) DeleteSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{uuid}", uuid, string(deleteMode))
}
