// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnregisterLicenseServer 操作UnregisterLicenseServer
func (cli *ZSClient) UnregisterLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/client", uuid, string(deleteMode))
}

