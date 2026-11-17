// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdatePciDevice updates PciDevice
func (cli *ZSClient) UpdatePciDevice(uuid string, params param.UpdatePciDeviceParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.PutWithRespKey("v1/pci-device/pci-devices", uuid, "", map[string]interface{}{
		"updatePciDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePciDevice deletes PciDevice
func (cli *ZSClient) DeletePciDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-devices", uuid, string(deleteMode))
}
// QueryPciDevice queries PciDevice list
func (cli *ZSClient) QueryPciDevice(params *param.QueryParam) ([]view.PciDeviceInventoryView, error) {
	var resp []view.PciDeviceInventoryView
	return resp, cli.List("v1/pci-device/pci-devices", params, &resp)
}

func (cli *ZSClient) GetPciDevice(uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.Get("v1/pci-device/pci-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDevice Pagination
func (cli *ZSClient) PagePciDevice(params *param.QueryParam) ([]view.PciDeviceInventoryView, int, error) {
	var pciDevices []view.PciDeviceInventoryView
	total, err := cli.Page("v1/pci-device/pci-devices", params, &pciDevices)
	return pciDevices, total, err
}

// UngenerateSriovPciDevices operates on UngenerateSriovPciDevices
func (cli *ZSClient) UngenerateSriovPciDevices(pciDeviceUuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey("v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"ungenerateSriovPciDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPciDeviceToVm operates on PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(pciDeviceUuid string, params param.AttachPciDeviceToVmParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/pci-device/pci-devices/%s/attach", pciDeviceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForAttachingVm gets PciDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForAttachingVm(uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPciDeviceFromVm operates on PciDeviceFromVm
func (cli *ZSClient) DetachPciDeviceFromVm(pciDeviceUuid string, params param.DetachPciDeviceFromVmParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/pci-device/pci-devices/%s/detach", pciDeviceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForNewCreateVm gets PciDeviceCandidatesForNewCreateVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForNewCreateVm() (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.GetWithRespKey("v1/pci-device/candidate-pci-devices-for-new-create-vm", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEthernetVF queries EthernetVF list
func (cli *ZSClient) QueryEthernetVF(params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, error) {
	var resp []view.EthernetVfPciDeviceInventoryView
	return resp, cli.List("v1/pci-device/ethernet-vfs", params, &resp)
}

func (cli *ZSClient) GetEthernetVF(uuid string) (*view.EthernetVfPciDeviceInventoryView, error) {
	var resp view.EthernetVfPciDeviceInventoryView
	if err := cli.Get("v1/pci-device/ethernet-vfs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEthernetVF Pagination
func (cli *ZSClient) PageEthernetVF(params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, int, error) {
	var ethernetVFs []view.EthernetVfPciDeviceInventoryView
	total, err := cli.Page("v1/pci-device/ethernet-vfs", params, &ethernetVFs)
	return ethernetVFs, total, err
}

// GenerateSriovPciDevices operates on SriovPciDevices
func (cli *ZSClient) GenerateSriovPciDevices(pciDeviceUuid string, params param.GenerateSriovPciDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithRespKey("v1/pci-devices", pciDeviceUuid, "", map[string]interface{}{
		"generateSriovPciDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
