// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPciDeviceSpec queries PciDeviceSpec list
func (cli *ZSClient) QueryPciDeviceSpec(params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, error) {
	var resp []view.PciDeviceSpecInventoryView
	return resp, cli.List("v1/pci-device-specs", params, &resp)
}

func (cli *ZSClient) GetPciDeviceSpec(uuid string) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.PciDeviceSpecInventoryView
	if err := cli.Get("v1/pci-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDeviceSpec Pagination
func (cli *ZSClient) PagePciDeviceSpec(params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, int, error) {
	var pciDeviceSpecs []view.PciDeviceSpecInventoryView
	total, err := cli.Page("v1/pci-device-specs", params, &pciDeviceSpecs)
	return pciDeviceSpecs, total, err
}
// UpdatePciDeviceSpec updates PciDeviceSpec
func (cli *ZSClient) UpdatePciDeviceSpec(uuid string, params param.UpdatePciDeviceSpecParam) (*view.PciDeviceSpecInventoryView, error) {
	resp := view.PciDeviceSpecInventoryView{}
	if err := cli.PutWithRespKey("v1/pci-device-specs", uuid, "", map[string]interface{}{
		"updatePciDeviceSpec": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePciDeviceSpecFromVmInstance removes PciDeviceSpecFromVmInstance
func (cli *ZSClient) RemovePciDeviceSpecFromVmInstance(pciSpecUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/pci-device-specs", pciSpecUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddPciDeviceSpecToVmInstance adds PciDeviceSpecToVmInstance
func (cli *ZSClient) AddPciDeviceSpecToVmInstance(pciSpecUuid string, vmInstanceUuid string, params param.AddPciDeviceSpecToVmInstanceParam) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	resp := view.VmInstancePciDeviceSpecRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/pci-device-specs/%s/vm-instances/%s", pciSpecUuid, vmInstanceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceSpecCandidates gets PciDeviceSpecCandidates by uuid
func (cli *ZSClient) GetPciDeviceSpecCandidates() (*view.PciDeviceSpecInventoryView, error) {
	var resp view.PciDeviceSpecInventoryView
	if err := cli.GetWithRespKey("v1/pci-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
