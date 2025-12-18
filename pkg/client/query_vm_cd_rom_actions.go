// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmCdRom queries VmCdRom list
func (cli *ZSClient) QueryVmCdRom(params param.QueryParam) ([]view.VmCdRomInventoryView, error) {
	var resp []view.VmCdRomInventoryView
	return resp, cli.List("v1/vm-instances/cdroms", &params, &resp)
}
