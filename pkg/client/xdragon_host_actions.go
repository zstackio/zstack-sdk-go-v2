// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddXDragonHost adds XDragonHost
func (cli *ZSClient) AddXDragonHost(params param.AddXDragonHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.Post("v1/hosts/xdragon", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
