// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBareMetal2ChassisOffering updates BareMetal2ChassisOffering
func (cli *ZSClient) UpdateBareMetal2ChassisOffering(uuid string, params param.UpdateBareMetal2ChassisOfferingParam) (*view.BareMetal2ChassisOfferingInventoryView, error) {
	var resp view.UpdateBareMetal2ChassisOfferingEventView
	err := cli.PutWithSpec("v1/baremetal2/chassis/offerings", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBareMetal2ChassisOffering queries BareMetal2ChassisOffering list
func (cli *ZSClient) QueryBareMetal2ChassisOffering(params *param.QueryParam) ([]view.BareMetal2ChassisOfferingInventoryView, error) {
	var resp []view.BareMetal2ChassisOfferingInventoryView
	return resp, cli.List("v1/baremetal2/chassis/offerings", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ChassisOffering(uuid string) (*view.BareMetal2ChassisOfferingInventoryView, error) {
	var resp view.BareMetal2ChassisOfferingInventoryView
	if err := cli.Get("v1/baremetal2/chassis/offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
