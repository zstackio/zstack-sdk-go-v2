// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateL2GeneveNetwork creates L2GeneveNetwork
func (cli *ZSClient) CreateL2GeneveNetwork(ctx context.Context, params param.CreateL2GeneveNetworkParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/l2-networks/geneve", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
