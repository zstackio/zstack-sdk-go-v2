// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHybridEipFromLocal 删除HybridEipFromLocal
func (cli *ZSClient) DeleteHybridEipFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/eip/{uuid}", uuid, string(deleteMode))
}

