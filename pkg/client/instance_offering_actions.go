// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeInstanceOffering changes InstanceOffering
func (cli *ZSClient) ChangeInstanceOffering(uuid string, params param.ChangeInstanceOfferingParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ChangeInstanceOfferingEventView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateInstanceOffering updates InstanceOffering
func (cli *ZSClient) UpdateInstanceOffering(uuid string, params param.UpdateInstanceOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.UpdateInstanceOfferingEventView
	if err := cli.Put("v1/instance-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateInstanceOffering creates InstanceOffering
func (cli *ZSClient) CreateInstanceOffering(params param.CreateInstanceOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryInstanceOffering queries InstanceOffering list
func (cli *ZSClient) QueryInstanceOffering(params *param.QueryParam) ([]view.InstanceOfferingInventoryView, error) {
	var resp []view.InstanceOfferingInventoryView
	return resp, cli.List("v1/instance-offerings", params, &resp)
}

func (cli *ZSClient) GetInstanceOffering(uuid string) (*view.InstanceOfferingInventoryView, error) {
	var resp view.InstanceOfferingInventoryView
	if err := cli.Get("v1/instance-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteInstanceOffering deletes InstanceOffering
func (cli *ZSClient) DeleteInstanceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/instance-offerings", uuid, string(deleteMode))
}
