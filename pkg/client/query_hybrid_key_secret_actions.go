// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHybridKeySecret queries HybridKeySecret list
func (cli *ZSClient) QueryHybridKeySecret(params *param.QueryParam) ([]view.HybridAccountInventoryView, error) {
	var resp []view.HybridAccountInventoryView
	return resp, cli.List("v1/hybrid/hybrid/key", params, &resp)
}
