// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// CleanupBillingUsage operates on CleanupBillingUsage
func (cli *ZSClient) CleanupBillingUsage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/usage", uuid, string(deleteMode))
}
