// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterDatacenter queries VCenterDatacenter list
func (cli *ZSClient) QueryVCenterDatacenter(params *param.QueryParam) ([]view.VCenterDatacenterInventoryView, error) {
	var resp []view.VCenterDatacenterInventoryView
	return resp, cli.List("v1/vcenters/datacenters", params, &resp)
}

func (cli *ZSClient) GetVCenterDatacenter(uuid string) (*view.VCenterDatacenterInventoryView, error) {
	var resp view.VCenterDatacenterInventoryView
	if err := cli.Get("v1/vcenters/datacenters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterDatacenter Pagination
func (cli *ZSClient) PageVCenterDatacenter(params *param.QueryParam) ([]view.VCenterDatacenterInventoryView, int, error) {
	var vCenterDatacenters []view.VCenterDatacenterInventoryView
	total, err := cli.Page("v1/vcenters/datacenters", params, &vCenterDatacenters)
	return vCenterDatacenters, total, err
}
