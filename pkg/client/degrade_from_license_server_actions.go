// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DegradeFromLicenseServer operates on DegradeFromLicenseServer
func (cli *ZSClient) DegradeFromLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server", uuid, string(deleteMode))
}
