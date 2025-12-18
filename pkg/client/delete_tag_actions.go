// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteTag deletes Tag
func (cli *ZSClient) DeleteTag(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags/{uuid}", uuid, string(deleteMode))
}
