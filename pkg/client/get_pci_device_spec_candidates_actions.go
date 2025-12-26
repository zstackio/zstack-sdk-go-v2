// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPciDeviceSpecCandidates gets PciDeviceSpecCandidates by uuid
func (cli *ZSClient) GetPciDeviceSpecCandidates(uuid string) (*view.GetPciDeviceSpecCandidatesView, error) {
	var resp view.GetPciDeviceSpecCandidatesView
	if err := cli.Get("v1/pci-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
