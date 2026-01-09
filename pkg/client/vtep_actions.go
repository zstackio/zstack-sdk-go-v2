// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVtep queries Vtep list
func (cli *ZSClient) QueryVtep(params *param.QueryParam) ([]view.VtepInventoryView, error) {
	var resp []view.VtepInventoryView
	return resp, cli.List("v1/l2-networks/vteps", params, &resp)
}

func (cli *ZSClient) GetVtep(uuid string) (*view.VtepInventoryView, error) {
	var resp view.VtepInventoryView
	if err := cli.Get("v1/l2-networks/vteps", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
