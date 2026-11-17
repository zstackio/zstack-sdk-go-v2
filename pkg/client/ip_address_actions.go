// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteIpAddress deletes IpAddress
func (cli *ZSClient) DeleteIpAddress(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks", uuid, string(deleteMode))
}

// QueryIpAddress queries IpAddress list
func (cli *ZSClient) QueryIpAddress(params *param.QueryParam) ([]view.UsedIpInventoryView, error) {
	var resp []view.UsedIpInventoryView
	return resp, cli.List("v1/l3-networks/ip-address", params, &resp)
}

func (cli *ZSClient) GetIpAddress(uuid string) (*view.UsedIpInventoryView, error) {
	var resp view.UsedIpInventoryView
	if err := cli.Get("v1/l3-networks/ip-address", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIpAddress Pagination
func (cli *ZSClient) PageIpAddress(params *param.QueryParam) ([]view.UsedIpInventoryView, int, error) {
	var ipAddress []view.UsedIpInventoryView
	total, err := cli.Page("v1/l3-networks/ip-address", params, &ipAddress)
	return ipAddress, total, err
}
