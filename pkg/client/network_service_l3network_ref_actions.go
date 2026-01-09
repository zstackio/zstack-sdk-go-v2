// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNetworkServiceL3NetworkRef queries NetworkServiceL3NetworkRef list
func (cli *ZSClient) QueryNetworkServiceL3NetworkRef(params *param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp []view.NetworkServiceL3NetworkRefInventoryView
	return resp, cli.List("v1/l3-networks/network-services/refs", params, &resp)
}

func (cli *ZSClient) GetNetworkServiceL3NetworkRef(uuid string) (*view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp view.NetworkServiceL3NetworkRefInventoryView
	if err := cli.Get("v1/l3-networks/network-services/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
