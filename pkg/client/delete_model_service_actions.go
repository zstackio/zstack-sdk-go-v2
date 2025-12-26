// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteModelService deletes ModelService
func (cli *ZSClient) DeleteModelService(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-services/{uuid}", uuid, string(deleteMode))
}
