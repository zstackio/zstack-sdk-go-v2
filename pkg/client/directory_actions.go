// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteDirectory 删除Directory
func (cli *ZSClient) DeleteDirectory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/directory", uuid, string(deleteMode))
}

