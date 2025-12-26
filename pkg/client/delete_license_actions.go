// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/licenses/mn/{managementNodeUuid}/actions", uuid, string(deleteMode))
}
