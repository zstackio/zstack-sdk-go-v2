// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHybridKeySecret 删除HybridKeySecret
func (cli *ZSClient) DeleteHybridKeySecret(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/hybrid/key/{uuid}", uuid, string(deleteMode))
}

