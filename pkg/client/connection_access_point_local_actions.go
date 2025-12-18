// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteConnectionAccessPointLocal 删除ConnectionAccessPointLocal
func (cli *ZSClient) DeleteConnectionAccessPointLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/access-point/{uuid}", uuid, string(deleteMode))
}

