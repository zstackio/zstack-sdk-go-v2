// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHostPhysicalMemory queries HostPhysicalMemory list
func (cli *ZSClient) QueryHostPhysicalMemory(params *param.QueryParam) ([]view.HostPhysicalMemoryInventoryView, error) {
	var resp []view.HostPhysicalMemoryInventoryView
	return resp, cli.List("v1/hosts/physicalmemory", params, &resp)
}

func (cli *ZSClient) GetHostPhysicalMemory(uuid string) (*view.HostPhysicalMemoryInventoryView, error) {
	var resp view.HostPhysicalMemoryInventoryView
	if err := cli.Get("v1/hosts/physicalmemory", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
