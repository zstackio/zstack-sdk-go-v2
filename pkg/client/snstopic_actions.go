// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSNSTopic 删除SNSTopic
func (cli *ZSClient) DeleteSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{uuid}", uuid, string(deleteMode))
}

