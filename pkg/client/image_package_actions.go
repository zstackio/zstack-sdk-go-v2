// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteImagePackage 删除ImagePackage
func (cli *ZSClient) DeleteImagePackage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-packages/{uuid}", uuid, string(deleteMode))
}

