// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteModelCenter deletes ModelCenter
func (cli *ZSClient) DeleteModelCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-centers/{uuid}", uuid, string(deleteMode))
}
