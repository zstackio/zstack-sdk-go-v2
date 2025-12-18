// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSdnController 操作RemoveSdnController
func (cli *ZSClient) RemoveSdnController(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controllers/{uuid}", uuid, string(deleteMode))
}

