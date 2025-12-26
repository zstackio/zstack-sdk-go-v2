// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePriceTable deletes PriceTable
func (cli *ZSClient) DeletePriceTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/price-tables/{uuid}", uuid, string(deleteMode))
}
