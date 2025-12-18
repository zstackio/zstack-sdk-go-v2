// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/licenses/mn/{managementNodeUuid}/actions", uuid, string(deleteMode))
}
