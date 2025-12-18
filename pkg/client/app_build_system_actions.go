// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAppBuildSystem 删除AppBuildSystem
func (cli *ZSClient) DeleteAppBuildSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildsystem/{uuid}", uuid, string(deleteMode))
}

