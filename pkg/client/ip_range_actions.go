// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddIpRange adds IpRange
func (cli *ZSClient) AddIpRange(l3NetworkUuid string, params param.AddIpRangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/ip-ranges", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIpRange updates IpRange
func (cli *ZSClient) UpdateIpRange(uuid string, params param.UpdateIpRangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PutWithRespKey("v1/l3-networks/ip-ranges", uuid, "", map[string]interface{}{
		"updateIpRange": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIpRange queries IpRange list
func (cli *ZSClient) QueryIpRange(params *param.QueryParam) ([]view.IpRangeInventoryView, error) {
	var resp []view.IpRangeInventoryView
	return resp, cli.List("v1/l3-networks/ip-ranges", params, &resp)
}

func (cli *ZSClient) GetIpRange(uuid string) (*view.IpRangeInventoryView, error) {
	var resp view.IpRangeInventoryView
	if err := cli.Get("v1/l3-networks/ip-ranges", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIpRange Pagination
func (cli *ZSClient) PageIpRange(params *param.QueryParam) ([]view.IpRangeInventoryView, int, error) {
	var ipRanges []view.IpRangeInventoryView
	total, err := cli.Page("v1/l3-networks/ip-ranges", params, &ipRanges)
	return ipRanges, total, err
}
// DeleteIpRange deletes IpRange
func (cli *ZSClient) DeleteIpRange(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks/ip-ranges", uuid, string(deleteMode))
}
