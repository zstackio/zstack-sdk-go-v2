// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNfvInstOffering queries NfvInstOffering list
func (cli *ZSClient) QueryNfvInstOffering(params *param.QueryParam) ([]view.NfvInstOfferingInventoryView, error) {
	var resp []view.NfvInstOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/nfvinst", params, &resp)
}
