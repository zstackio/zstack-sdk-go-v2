// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DegradeFromLicenseServer 操作DegradeFromLicenseServer
func (cli *ZSClient) DegradeFromLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server", uuid, string(deleteMode))
}

