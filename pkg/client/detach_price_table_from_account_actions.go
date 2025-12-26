// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPriceTableFromAccount operates on PriceTableFromAccount
func (cli *ZSClient) DetachPriceTableFromAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/price-tables/{tableUuid}/accounts/{accountUuid}", uuid, string(deleteMode))
}
