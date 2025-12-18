// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteExternalBackup deletes ExternalBackup
func (cli *ZSClient) DeleteExternalBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/externalbackup/{uuid}", uuid, string(deleteMode))
}
