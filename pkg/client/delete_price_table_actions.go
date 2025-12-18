// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePriceTable deletes PriceTable
func (cli *ZSClient) DeletePriceTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/price-tables/{uuid}", uuid, string(deleteMode))
}
