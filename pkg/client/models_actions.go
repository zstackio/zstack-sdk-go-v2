// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteModels 删除Models
func (cli *ZSClient) DeleteModels(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models", uuid, string(deleteMode))
}

