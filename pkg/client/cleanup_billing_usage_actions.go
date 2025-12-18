// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanupBillingUsage 操作CleanupBillingUsage
func (cli *ZSClient) CleanupBillingUsage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/usage", uuid, string(deleteMode))
}

