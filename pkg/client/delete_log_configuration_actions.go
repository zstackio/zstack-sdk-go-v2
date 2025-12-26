// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteLogConfiguration deletes LogConfiguration
func (cli *ZSClient) DeleteLogConfiguration(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/log/configurations/log4j2", uuid, string(deleteMode))
}
