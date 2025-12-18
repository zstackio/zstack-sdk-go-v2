// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteModel deletes Model
func (cli *ZSClient) DeleteModel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{uuid}", uuid, string(deleteMode))
}
