// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteNasFileSystem deletes NasFileSystem
func (cli *ZSClient) DeleteNasFileSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/{uuid}", uuid, string(deleteMode))
}
