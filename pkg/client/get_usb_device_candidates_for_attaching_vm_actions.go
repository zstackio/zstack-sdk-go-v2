// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetUsbDeviceCandidatesForAttachingVm gets UsbDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetUsbDeviceCandidatesForAttachingVm(uuid string) (*view.GetUsbDeviceCandidatesForAttachingVmView, error) {
	var resp view.GetUsbDeviceCandidatesForAttachingVmView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/candidate-usb-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
