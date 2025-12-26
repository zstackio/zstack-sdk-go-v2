// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModelServices deletes ModelServices
func (cli *ZSClient) DeleteModelServices(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/", uuid, string(deleteMode))
}
