// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHost 删除Host
func (cli *ZSClient) DeleteHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/{uuid}", uuid, string(deleteMode))
}

