// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHybridEipFromLocal queries HybridEipFromLocal list
func (cli *ZSClient) QueryHybridEipFromLocal(params *param.QueryParam) ([]view.HybridEipAddressInventoryView, error) {
	var resp []view.HybridEipAddressInventoryView
	return resp, cli.List("v1/hybrid/eip", params, &resp)
}
