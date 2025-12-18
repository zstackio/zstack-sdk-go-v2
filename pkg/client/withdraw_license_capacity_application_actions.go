// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// WithdrawLicenseCapacityApplication operates on WithdrawLicenseCapacityApplication
func (cli *ZSClient) WithdrawLicenseCapacityApplication(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/capacity-application", uuid, string(deleteMode))
}
