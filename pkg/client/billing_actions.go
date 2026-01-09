// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBilling deletes Billing
func (cli *ZSClient) DeleteBilling(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/billings", uuid, string(deleteMode))
}
