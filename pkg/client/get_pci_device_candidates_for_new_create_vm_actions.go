// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPciDeviceCandidatesForNewCreateVm gets PciDeviceCandidatesForNewCreateVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForNewCreateVm(uuid string) (*view.GetPciDeviceCandidatesForNewCreateVmView, error) {
	var resp view.GetPciDeviceCandidatesForNewCreateVmView
	if err := cli.Get("v1/pci-device/candidate-pci-devices-for-new-create-vm", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
