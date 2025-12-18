// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteEcsVpcInLocal 删除EcsVpcInLocal
func (cli *ZSClient) DeleteEcsVpcInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc/{uuid}", uuid, string(deleteMode))
}

