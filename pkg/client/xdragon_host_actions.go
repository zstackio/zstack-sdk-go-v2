// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddXDragonHost adds XDragonHost
func (cli *ZSClient) AddXDragonHost(ctx context.Context, params param.AddXDragonHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/hosts/xdragon", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
