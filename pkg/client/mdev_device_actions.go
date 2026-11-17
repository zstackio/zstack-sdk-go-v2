// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDevice updates MdevDevice
func (cli *ZSClient) UpdateMdevDevice(uuid string, params param.UpdateMdevDeviceParam) (*view.MdevDeviceInventoryView, error) {
	resp := view.MdevDeviceInventoryView{}
	if err := cli.PutWithRespKey("v1/mdev-devices", uuid, "", map[string]interface{}{
		"updateMdevDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMdevDevice deletes MdevDevice
func (cli *ZSClient) DeleteMdevDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-devices", uuid, string(deleteMode))
}
// QueryMdevDevice queries MdevDevice list
func (cli *ZSClient) QueryMdevDevice(params *param.QueryParam) ([]view.MdevDeviceInventoryView, error) {
	var resp []view.MdevDeviceInventoryView
	return resp, cli.List("v1/mdev-devices", params, &resp)
}

func (cli *ZSClient) GetMdevDevice(uuid string) (*view.MdevDeviceInventoryView, error) {
	var resp view.MdevDeviceInventoryView
	if err := cli.Get("v1/mdev-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMdevDevice Pagination
func (cli *ZSClient) PageMdevDevice(params *param.QueryParam) ([]view.MdevDeviceInventoryView, int, error) {
	var mdevDevices []view.MdevDeviceInventoryView
	total, err := cli.Page("v1/mdev-devices", params, &mdevDevices)
	return mdevDevices, total, err
}

// GetMdevDeviceCandidates gets MdevDeviceCandidates by uuid
func (cli *ZSClient) GetMdevDeviceCandidates() (*view.MdevDeviceInventoryView, error) {
	var resp view.MdevDeviceInventoryView
	if err := cli.GetWithRespKey("v1/mdev-devices/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateHygonMdevDevices operates on HygonMdevDevices
func (cli *ZSClient) GenerateHygonMdevDevices(hostUuid string, params param.GenerateHygonMdevDevicesParam) (*view.GenerateHygonMdevDevicesEventView, error) {
	resp := view.GenerateHygonMdevDevicesEventView{}
	if err := cli.PutWithRespKey("v1/hygon-devices", hostUuid, "", map[string]interface{}{
		"generateHygonMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachMdevDeviceToVm operates on MdevDeviceToVm
func (cli *ZSClient) AttachMdevDeviceToVm(mdevDeviceUuid string, vmInstanceUuid string, params param.AttachMdevDeviceToVmParam) (*view.MdevDeviceInventoryView, error) {
	resp := view.MdevDeviceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/mdev-devices/%s/vm-instances/%s", mdevDeviceUuid, vmInstanceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateHygonMdevDevices operates on UngenerateHygonMdevDevices
func (cli *ZSClient) UngenerateHygonMdevDevices(hostUuid string, params param.UngenerateHygonMdevDevicesParam) (*view.UngenerateHygonMdevDevicesEventView, error) {
	resp := view.UngenerateHygonMdevDevicesEventView{}
	if err := cli.PutWithRespKey("v1/hygon-devices", hostUuid, "", map[string]interface{}{
		"ungenerateHygonMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachMdevDeviceFromVm operates on MdevDeviceFromVm
func (cli *ZSClient) DetachMdevDeviceFromVm(mdevDeviceUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/mdev-devices", mdevDeviceUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GenerateMdevDevices operates on MdevDevices
func (cli *ZSClient) GenerateMdevDevices(pciDeviceUuid string, params param.GenerateMdevDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey("v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"generateMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateMdevDevices operates on UngenerateMdevDevices
func (cli *ZSClient) UngenerateMdevDevices(pciDeviceUuid string, params param.UngenerateMdevDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey("v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"ungenerateMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
