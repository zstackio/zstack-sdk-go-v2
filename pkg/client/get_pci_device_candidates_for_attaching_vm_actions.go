// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPciDeviceCandidatesForAttachingVm gets PciDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForAttachingVm(uuid string) (*view.GetPciDeviceCandidatesForAttachingVmView, error) {
	var resp view.GetPciDeviceCandidatesForAttachingVmView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-pci-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
