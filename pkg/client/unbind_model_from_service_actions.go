// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// UnbindModelFromService operates on UnbindModelFromService
func (cli *ZSClient) UnbindModelFromService(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", uuid, string(deleteMode))
}
