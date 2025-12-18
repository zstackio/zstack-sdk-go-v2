// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteMedia 删除Media
func (cli *ZSClient) DeleteMedia(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/media/{uuid}", uuid, string(deleteMode))
}

