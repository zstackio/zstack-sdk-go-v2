// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPciDeviceOffering queries PciDeviceOffering list
func (cli *ZSClient) QueryPciDeviceOffering(params param.QueryParam) ([]view.PciDeviceOfferingInventoryView, error) {
	var resp []view.PciDeviceOfferingInventoryView
	return resp, cli.List("v1/pci-device/pci-device-offerings", &params, &resp)
}
