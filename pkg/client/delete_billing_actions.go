// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBilling deletes Billing
func (cli *ZSClient) DeleteBilling(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/billings", uuid, string(deleteMode))
}
