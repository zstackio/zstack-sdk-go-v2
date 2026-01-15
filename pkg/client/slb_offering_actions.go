// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSlbOffering creates SlbOffering
func (cli *ZSClient) CreateSlbOffering(params param.CreateSlbOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post("v1/instance-offerings/slb", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySlbOffering queries SlbOffering list
func (cli *ZSClient) QuerySlbOffering(params *param.QueryParam) ([]view.SlbOfferingInventoryView, error) {
	var resp []view.SlbOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/slb", params, &resp)
}

// PageSlbOffering Pagination
func (cli *ZSClient) PageSlbOffering(params *param.QueryParam) ([]view.SlbOfferingInventoryView, int, error) {
	var slbOfferings []view.SlbOfferingInventoryView
	total, err := cli.Page("v1/instance-offerings/slb", params, &slbOfferings)
	return slbOfferings, total, err
}
