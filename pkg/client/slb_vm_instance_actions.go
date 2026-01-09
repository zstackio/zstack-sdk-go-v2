// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySlbVmInstance queries SlbVmInstance list
func (cli *ZSClient) QuerySlbVmInstance(params *param.QueryParam) ([]view.SlbVmInstanceInventoryView, error) {
	var resp []view.SlbVmInstanceInventoryView
	return resp, cli.List("v1/load-balancers/slb/instances", params, &resp)
}

func (cli *ZSClient) GetSlbVmInstance(uuid string) (*view.SlbVmInstanceInventoryView, error) {
	var resp view.SlbVmInstanceInventoryView
	if err := cli.Get("v1/load-balancers/slb/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
