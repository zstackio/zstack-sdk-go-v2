// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeInstanceOffering changes InstanceOffering
func (cli *ZSClient) ChangeInstanceOffering(vmInstanceUuid string, params param.ChangeInstanceOfferingParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"changeInstanceOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateInstanceOffering updates InstanceOffering
func (cli *ZSClient) UpdateInstanceOffering(uuid string, params param.UpdateInstanceOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithRespKey("v1/instance-offerings", uuid, "", map[string]interface{}{
		"updateInstanceOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateInstanceOffering creates InstanceOffering
func (cli *ZSClient) CreateInstanceOffering(params param.CreateInstanceOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post("v1/instance-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageInstanceOffering Pagination
func (cli *ZSClient) PageInstanceOffering(params *param.QueryParam) ([]view.InstanceOfferingInventoryView, int, error) {
	var instanceOfferings []view.InstanceOfferingInventoryView
	total, err := cli.Page("v1/instance-offerings", params, &instanceOfferings)
	return instanceOfferings, total, err
}
// DeleteInstanceOffering deletes InstanceOffering
func (cli *ZSClient) DeleteInstanceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/instance-offerings", uuid, string(deleteMode))
}

// CreateOvnControllerOffering creates OvnControllerOffering
func (cli *ZSClient) CreateOvnControllerOffering(params param.CreateOvnControllerOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post("v1/instance-offerings/ovn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeInstanceOfferingState changes InstanceOfferingState
func (cli *ZSClient) ChangeInstanceOfferingState(uuid string, params param.ChangeInstanceOfferingStateParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithRespKey("v1/instance-offerings", uuid, "", map[string]interface{}{
		"changeInstanceOfferingState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateInstanceOfferingUserConfig operates on InstanceOfferingUserConfig
func (cli *ZSClient) ValidateInstanceOfferingUserConfig(params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.PutWithRespKey("v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validateInstanceOfferingUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
