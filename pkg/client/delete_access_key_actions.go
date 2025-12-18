// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteAccessKey deletes AccessKey
func (cli *ZSClient) DeleteAccessKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accesskeys/{uuid}", uuid, string(deleteMode))
}
