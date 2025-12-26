// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// WithdrawLicenseCapacityApplication operates on WithdrawLicenseCapacityApplication
func (cli *ZSClient) WithdrawLicenseCapacityApplication(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/capacity-application", uuid, string(deleteMode))
}
