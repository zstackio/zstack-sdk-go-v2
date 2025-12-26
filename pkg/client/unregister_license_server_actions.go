// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// UnregisterLicenseServer operates on UnregisterLicenseServer
func (cli *ZSClient) UnregisterLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/client", uuid, string(deleteMode))
}
