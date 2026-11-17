// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryHygonDevice queries HygonDevice list
func (cli *ZSClient) QueryHygonDevice(params *param.QueryParam) ([]view.HygonCcpDeviceInventoryView, error) {
	var resp []view.HygonCcpDeviceInventoryView
	return resp, cli.List("v1/hygon-devices", params, &resp)
}

func (cli *ZSClient) GetHygonDevice(uuid string) (*view.HygonCcpDeviceInventoryView, error) {
	var resp view.HygonCcpDeviceInventoryView
	if err := cli.Get("v1/hygon-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHygonDevice Pagination
func (cli *ZSClient) PageHygonDevice(params *param.QueryParam) ([]view.HygonCcpDeviceInventoryView, int, error) {
	var hygonDevices []view.HygonCcpDeviceInventoryView
	total, err := cli.Page("v1/hygon-devices", params, &hygonDevices)
	return hygonDevices, total, err
}
