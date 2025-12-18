// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteCbtTask 删除CbtTask
func (cli *ZSClient) DeleteCbtTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cbt-task/{uuid}", uuid, string(deleteMode))
}

