// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePreconfigurationTemplate deletes PreconfigurationTemplate
func (cli *ZSClient) DeletePreconfigurationTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/preconfigurations/{uuid}", uuid, string(deleteMode))
}
