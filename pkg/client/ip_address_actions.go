// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteIpAddress requires additional arguments.
//
// Deprecated: use DeleteIpAddressWithParams to supply the required fields.
func (cli *ZSClient) DeleteIpAddress(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteIpAddress requires additional arguments; use DeleteIpAddressWithParams")
}

// DeleteIpAddressWithParams sends the required route and query parameters for DeleteIpAddress.
func (cli *ZSClient) DeleteIpAddressWithParams(l3NetworkUuid string, deleteMode param.DeleteMode, params param.DeleteIpAddressParamDetail) error {
	query := url.Values{
		"usedIpUuids": params.UsedIpUuids,
	}
	return cli.deleteActionWithQuery("v1/l3-networks", l3NetworkUuid, "ip-address", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "usedIpUuids")
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
