// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPciDeviceSpecCandidates 获取PciDeviceSpecCandidates详情
func (cli *ZSClient) GetPciDeviceSpecCandidates(uuid string) (*view.GetPciDeviceSpecCandidatesView, error) {
	var resp view.GetPciDeviceSpecCandidatesView
	if err := cli.Get("v1/pci-device-specs/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

