// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteBilling deletes Billing
func (cli *ZSClient) DeleteBilling(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/billings", uuid, string(deleteMode))
}
