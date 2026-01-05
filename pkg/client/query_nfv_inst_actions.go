// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNfvInst queries NfvInst list
func (cli *ZSClient) QueryNfvInst(params *param.QueryParam) ([]view.NfvInstInventoryView, error) {
	var resp []view.NfvInstInventoryView
	return resp, cli.List("v1/vm-instances/appliances/nfvinst", params, &resp)
}
