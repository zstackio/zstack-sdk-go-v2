// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnbindModelFromService 操作UnbindModelFromService
func (cli *ZSClient) UnbindModelFromService(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", uuid, string(deleteMode))
}

