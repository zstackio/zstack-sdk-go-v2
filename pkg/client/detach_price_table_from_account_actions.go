// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPriceTableFromAccount operates on PriceTableFromAccount
func (cli *ZSClient) DetachPriceTableFromAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", uuid, string(deleteMode))
}
