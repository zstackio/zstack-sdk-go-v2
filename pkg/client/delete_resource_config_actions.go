// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteResourceConfig deletes ResourceConfig
func (cli *ZSClient) DeleteResourceConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/resource-configurations/{category}/{name}/{resourceUuid}", uuid, string(deleteMode))
}
